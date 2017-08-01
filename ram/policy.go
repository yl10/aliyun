package ram

const (
	POLICY_VERSION      = "1"
	POLICY_EFFECT_ALLOW = "Allow"
	POLICY_EFFECT_DENY  = "Deny"
)

type EffectBlock string

type ActionString string
type ActionStrings []string
type ActionBlock interface{}

type ResourceString string
type ResourceStrings []string
type ResourceBlock interface{}

type ConditionValueString string
type ConditionValueNumber int
type ConditionValueBoolean bool
type ConditionMap map[string][]interface{}
type ConditionBlock map[string]ConditionMap

type StatementBlock struct {
	Effect      EffectBlock
	Action      ActionBlock    `json:"omitempty"`
	NotAction   ActionBlock    `json:"omitempty"`
	Resource    ResourceBlock  `json:"omitempty"`
	NotResource ResourceBlock  `json:"omitempty"`
	Condition   ConditionBlock `json:"omitempty"`
}

type Policy struct {
	Version   string
	Statement []StatementBlock
}

var (
	SimplePolicyTemplate = `
{
    "Version": "1",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "%s",
            "Resource": "%s"
        }
    ]
}`
)
