package basemodel

type Texts struct {
	Language string `json:"language" binding:"omitempty,language"`
	Key      string `json:"key" binding:"omitempty,textKey"`
	Value    string `json:"value"`
}

type TextsSlice []Texts

func (ts TextsSlice) Format() TextsSlice {
	unique := make(map[string]Texts)
	for _, t := range ts {
		// Формируем составной ключ из языка и ключа
		compositeKey := t.Language + "_" + t.Key
		if _, exists := unique[compositeKey]; !exists {
			unique[compositeKey] = t
		}
	}

	formatted := make(TextsSlice, 0, len(unique))
	for _, t := range unique {
		formatted = append(formatted, t)
	}
	return formatted
}
