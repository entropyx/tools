package operators

import pb "github.com/entropyx/protos"

func ToSqlOperator(operator pb.Operator) string {

	switch operator {
	case pb.Operator_EQUALS:
		return "="
	case pb.Operator_GREATER_THAN:
		return ">"
	case pb.Operator_LESS_THAN:
		return "<"
	}

	return ""
}
