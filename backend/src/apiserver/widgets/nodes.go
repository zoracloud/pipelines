package widgets

type Node interface {
	AddUpstreamNode()
	AddUpstreamNodes()
	RemoveUpstreamNode()
	AddDowntreamNode()
	AddDowntreamNodes()
	RemoveDownstreamNode()
}