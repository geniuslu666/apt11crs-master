package pkpass

// Field 定义通行证中的字段
type Field struct {
	Key                string                   `json:"key"`
	Label              string                   `json:"label,omitempty"`
	Value              interface{}              `json:"value"`
	AttributedValue    string                   `json:"attributedValue,omitempty"`
	ChangeMessage      string                   `json:"changeMessage,omitempty"`
	TextAlignment      string                   `json:"textAlignment,omitempty"`
	DataDetectorTypes  []string                 `json:"dataDetectorTypes,omitempty"`
	IsRelative         bool                     `json:"isRelative,omitempty"`
	DateStyle          string                   `json:"dateStyle,omitempty"`
	TimeStyle          string                   `json:"timeStyle,omitempty"`
	IgnoreTimeZone     bool                     `json:"ignoreTimeZone,omitempty"`
	CalendarIdentifier string                   `json:"calendarIdentifier,omitempty"`
	RelevantDate       string                   `json:"relevantDate,omitempty"`
	NumberStyle        string                   `json:"numberStyle,omitempty"`
	CurrencyCode       string                   `json:"currencyCode,omitempty"`
	Details            []map[string]interface{} `json:"details,omitempty"`
}

// TransitType 定义交通类型
type TransitType string

const (
	TransitTypeAir     TransitType = "PKTransitTypeAir"
	TransitTypeBoat    TransitType = "PKTransitTypeBoat"
	TransitTypeBus     TransitType = "PKTransitTypeBus"
	TransitTypeGeneric TransitType = "PKTransitTypeGeneric"
	TransitTypeTrain   TransitType = "PKTransitTypeTrain"
)

// PassStructure 定义通行证的结构
type PassStructure struct {
	PrimaryFields   []Field `json:"primaryFields,omitempty"`
	SecondaryFields []Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []Field `json:"auxiliaryFields,omitempty"`
	BackFields      []Field `json:"backFields,omitempty"`
	HeaderFields    []Field `json:"headerFields,omitempty"`
	TransitType     string  `json:"transitType,omitempty"`
}

// AddPrimaryField 添加主字段
func (p *Pass) AddPrimaryField(field Field) {
	if p.Data["primaryFields"] == nil {
		p.Data["primaryFields"] = []Field{}
	}

	fields := p.Data["primaryFields"].([]Field)
	fields = append(fields, field)
	p.Data["primaryFields"] = fields
}

// AddSecondaryField 添加次要字段
func (p *Pass) AddSecondaryField(field Field) {
	if p.Data["secondaryFields"] == nil {
		p.Data["secondaryFields"] = []Field{}
	}

	fields := p.Data["secondaryFields"].([]Field)
	fields = append(fields, field)
	p.Data["secondaryFields"] = fields
}

// AddAuxiliaryField 添加辅助字段
func (p *Pass) AddAuxiliaryField(field Field) {
	if p.Data["auxiliaryFields"] == nil {
		p.Data["auxiliaryFields"] = []Field{}
	}

	fields := p.Data["auxiliaryFields"].([]Field)
	fields = append(fields, field)
	p.Data["auxiliaryFields"] = fields
}

// AddBackField 添加背面字段
func (p *Pass) AddBackField(field Field) {
	if p.Data["backFields"] == nil {
		p.Data["backFields"] = []Field{}
	}

	fields := p.Data["backFields"].([]Field)
	fields = append(fields, field)
	p.Data["backFields"] = fields
}

// AddHeaderField 添加头部字段
func (p *Pass) AddHeaderField(field Field) {
	if p.Data["headerFields"] == nil {
		p.Data["headerFields"] = []Field{}
	}

	fields := p.Data["headerFields"].([]Field)
	fields = append(fields, field)
	p.Data["headerFields"] = fields
}

// SetTransitType 设置交通类型
func (p *Pass) SetTransitType(transitType TransitType) {
	p.Data["transitType"] = string(transitType)
}
