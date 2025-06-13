package util

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
)

//go:embed ser_deser_json_string.json
var templateFS embed.FS

// JsonTemplateSerDeserResolverUtil provides utility functions for resolving JSON templates
// from a predefined resource file with support for inheritance and references.
type JsonTemplateSerDeserResolverUtil struct {
	templatesRoot map[string]interface{}
	mutex         sync.RWMutex
	initialized   bool
}

// Template represents a single template definition
type Template struct {
	Content  interface{} `json:"content"`
	Inherits []string    `json:"inherits,omitempty"`
}

// Global instance
var globalResolver = &JsonTemplateSerDeserResolverUtil{}

const TEMPLATE_RESOURCE_PATH = "ser_deser_json_string.json"

// init initializes the global resolver instance
func init() {
	if err := globalResolver.loadTemplates(); err != nil {
		fmt.Printf("Failed to load templates: %v\n", err)
	}
}

// GetJsonString gets the JSON string for a specified template using the global resolver
func GetJsonString(templateName string) (string, error) {
	return globalResolver.GetJsonString(templateName)
}

// GetJsonString gets the JSON string for a specified template
func (j *JsonTemplateSerDeserResolverUtil) GetJsonString(templateName string) (string, error) {
	j.mutex.RLock()
	if !j.initialized {
		j.mutex.RUnlock()
		j.mutex.Lock()
		if !j.initialized {
			if err := j.loadTemplates(); err != nil {
				j.mutex.Unlock()
				return "", err
			}
		}
		j.mutex.Unlock()
		j.mutex.RLock()
	}
	j.mutex.RUnlock()

	// Get the template with inheritance handling
	processedTemplates := make(map[string]bool)
	resolvedNode, err := j.resolveTemplateWithInheritance(templateName, processedTemplates)
	if err != nil {
		return "", err
	}

	// Resolve references in the node
	processedDependencies := make(map[string]bool)
	if err := j.resolveReferences(resolvedNode, processedDependencies); err != nil {
		return "", err
	}

	// Convert to JSON string and return
	jsonBytes, err := json.Marshal(resolvedNode)
	if err != nil {
		return "", fmt.Errorf("failed to marshal resolved template: %w", err)
	}

	return string(jsonBytes), nil
}

// loadTemplates loads the templates from the predefined resource file
func (j *JsonTemplateSerDeserResolverUtil) loadTemplates() error {
	var reader io.Reader

	// Try to read from embedded filesystem first
	if file, err := templateFS.Open(TEMPLATE_RESOURCE_PATH); err == nil {
		defer file.Close()
		reader = file
	} else {
		return fmt.Errorf("resource not found: %s", TEMPLATE_RESOURCE_PATH)
	}

	var root map[string]interface{}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&root); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	templates, ok := root["templates"]
	if !ok {
		return fmt.Errorf("JSON template does not contain 'templates' root element")
	}

	templatesMap, ok := templates.(map[string]interface{})
	if !ok {
		return fmt.Errorf("'templates' element is not an object")
	}

	j.templatesRoot = templatesMap
	j.initialized = true
	return nil
}

// resolveTemplateWithInheritance resolves a template including all inherited fields from parent templates
func (j *JsonTemplateSerDeserResolverUtil) resolveTemplateWithInheritance(templateName string, processedTemplates map[string]bool) (interface{}, error) {
	if processedTemplates[templateName] {
		fmt.Printf("Warning: Circular inheritance detected for %s\n", templateName)
		return make(map[string]interface{}), nil
	}

	processedTemplates[templateName] = true

	templateData, ok := j.templatesRoot[templateName]
	if !ok {
		return nil, fmt.Errorf("template '%s' not found", templateName)
	}

	templateMap, ok := templateData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("template '%s' is not an object", templateName)
	}

	contentNode, ok := templateMap["content"]
	if !ok {
		return nil, fmt.Errorf("template '%s' does not contain 'content' node", templateName)
	}

	// If content is not an object, return it directly
	contentMap, isObject := contentNode.(map[string]interface{})
	if !isObject {
		return deepCopy(contentNode), nil
	}

	// Create a deep copy of the content node
	resultNode := deepCopy(contentMap).(map[string]interface{})

	// Process inheritance if present
	inheritsData, hasInherits := templateMap["inherits"]
	if hasInherits {
		inheritsSlice, ok := inheritsData.([]interface{})
		if ok {
			for _, parentNameInterface := range inheritsSlice {
				parentName, ok := parentNameInterface.(string)
				if !ok {
					continue
				}

				// Create a copy of processedTemplates for the parent resolution
				parentProcessedTemplates := make(map[string]bool)
				for k, v := range processedTemplates {
					parentProcessedTemplates[k] = v
				}

				// Resolve parent template
				parentNode, err := j.resolveTemplateWithInheritance(parentName, parentProcessedTemplates)
				if err != nil {
					return nil, err
				}

				// Only merge if parent is an object
				if parentMap, ok := parentNode.(map[string]interface{}); ok {
					j.mergeNodes(resultNode, parentMap)
				}
			}
		}
	}

	return resultNode, nil
}

// mergeNodes merges fields from the source node into the target node
// Fields in the target node are not overwritten if they already exist
func (j *JsonTemplateSerDeserResolverUtil) mergeNodes(target, source map[string]interface{}) {
	for fieldName, sourceValue := range source {
		// Only add the field if it doesn't exist in the target
		if _, exists := target[fieldName]; !exists {
			if targetMap, targetIsMap := target[fieldName].(map[string]interface{}); targetIsMap {
				if sourceMap, sourceIsMap := sourceValue.(map[string]interface{}); sourceIsMap {
					// Recursively merge objects
					j.mergeNodes(targetMap, sourceMap)
					continue
				}
			}
			// Add the field
			target[fieldName] = deepCopy(sourceValue)
		}
	}
}

// resolveReferences resolves references in a JSON node
func (j *JsonTemplateSerDeserResolverUtil) resolveReferences(node interface{}, processedDependencies map[string]bool) error {
	switch n := node.(type) {
	case map[string]interface{}:
		return j.resolveObjectReferences(n, processedDependencies)
	case []interface{}:
		return j.resolveArrayReferences(n, processedDependencies)
	}
	return nil
}

// resolveObjectReferences resolves references in an object node
func (j *JsonTemplateSerDeserResolverUtil) resolveObjectReferences(objectNode map[string]interface{}, processedDependencies map[string]bool) error {
	// Collect field names to avoid concurrent modification
	fieldsToProcess := make([]string, 0, len(objectNode))
	for fieldName := range objectNode {
		fieldsToProcess = append(fieldsToProcess, fieldName)
	}

	for _, fieldName := range fieldsToProcess {
		fieldValue, exists := objectNode[fieldName]
		if !exists {
			continue
		}

		// Check if the field name is a reference that needs to be resolved
		if j.isReference(fieldName) {
			referenceName := j.extractReferenceName(fieldName)

			// Create a copy of processed dependencies for this field
			fieldDependencies := make(map[string]bool)
			for k, v := range processedDependencies {
				fieldDependencies[k] = v
			}

			if fieldDependencies[referenceName] {
				fmt.Printf("Warning: Circular reference detected for %s\n", referenceName)
				continue
			}

			fieldDependencies[referenceName] = true

			// Resolve the template to get the actual key name
			resolvedReference, err := j.resolveTemplateWithInheritance(referenceName, make(map[string]bool))
			if err != nil {
				return err
			}

			// Only apply if the resolved reference is a simple value (string, number, etc.)
			if resolvedKey, ok := resolvedReference.(string); ok {
				// Remove the original reference key and add the resolved key with the same value
				originalValue := objectNode[fieldName]
				delete(objectNode, fieldName)
				objectNode[resolvedKey] = originalValue

				// Update the field name and value for further processing
				fieldName = resolvedKey
				fieldValue = originalValue
			}
		}

		// Check if the field value is a string reference
		if textValue, ok := fieldValue.(string); ok {
			if j.isReference(textValue) {
				referenceName := j.extractReferenceName(textValue)

				// Create a copy of processed dependencies for this field
				fieldDependencies := make(map[string]bool)
				for k, v := range processedDependencies {
					fieldDependencies[k] = v
				}

				if fieldDependencies[referenceName] {
					fmt.Printf("Warning: Circular reference detected for %s\n", referenceName)
					continue
				}

				fieldDependencies[referenceName] = true

				// Resolve the template WITH inheritance
				resolvedReference, err := j.resolveTemplateWithInheritance(referenceName, make(map[string]bool))
				if err != nil {
					return err
				}

				// Resolve any references in the resolved template
				if err := j.resolveReferences(resolvedReference, fieldDependencies); err != nil {
					return err
				}

				objectNode[fieldName] = resolvedReference
			}
		} else {
			// Handle nested objects and arrays
			switch v := fieldValue.(type) {
			case map[string]interface{}, []interface{}:
				// Create a copy of processed dependencies for nested structures
				nestedDependencies := make(map[string]bool)
				for k, v := range processedDependencies {
					nestedDependencies[k] = v
				}
				if err := j.resolveReferences(v, nestedDependencies); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// resolveArrayReferences resolves references in an array node
func (j *JsonTemplateSerDeserResolverUtil) resolveArrayReferences(arrayNode []interface{}, processedDependencies map[string]bool) error {
	for i, element := range arrayNode {
		if textValue, ok := element.(string); ok {
			if j.isReference(textValue) {
				referenceName := j.extractReferenceName(textValue)

				// Clone the dependencies for each array element
				elementDependencies := make(map[string]bool)
				for k, v := range processedDependencies {
					elementDependencies[k] = v
				}

				if elementDependencies[referenceName] {
					fmt.Printf("Warning: Circular reference detected for %s\n", referenceName)
					continue
				}

				elementDependencies[referenceName] = true

				// Resolve the template WITH inheritance
				resolvedReference, err := j.resolveTemplateWithInheritance(referenceName, make(map[string]bool))
				if err != nil {
					return err
				}

				// Resolve any references in the resolved template
				if err := j.resolveReferences(resolvedReference, elementDependencies); err != nil {
					return err
				}

				arrayNode[i] = resolvedReference
			}
		} else {
			// Recursively process nested objects and arrays
			switch v := element.(type) {
			case map[string]interface{}, []interface{}:
				if err := j.resolveReferences(v, make(map[string]bool)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// isReference checks if a string value is a template reference
func (j *JsonTemplateSerDeserResolverUtil) isReference(value string) bool {
	return strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}")
}

// extractReferenceName extracts the reference name from a reference string
func (j *JsonTemplateSerDeserResolverUtil) extractReferenceName(reference string) string {
	if len(reference) <= 3 {
		return ""
	}
	return reference[2 : len(reference)-1]
}

// deepCopy creates a deep copy of an interface{}
func deepCopy(src interface{}) interface{} {
	switch s := src.(type) {
	case map[string]interface{}:
		dst := make(map[string]interface{})
		for k, v := range s {
			dst[k] = deepCopy(v)
		}
		return dst
	case []interface{}:
		dst := make([]interface{}, len(s))
		for i, v := range s {
			dst[i] = deepCopy(v)
		}
		return dst
	default:
		return s
	}
}

// NewJsonTemplateSerDeserResolverUtil creates a new instance of the resolver
func NewJsonTemplateSerDeserResolverUtil() *JsonTemplateSerDeserResolverUtil {
	resolver := &JsonTemplateSerDeserResolverUtil{}
	if err := resolver.loadTemplates(); err != nil {
		fmt.Printf("Failed to load templates: %v\n", err)
	}
	return resolver
}
