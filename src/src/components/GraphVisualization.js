import React from 'react';
import Graph from 'react-graph-vis';

const GraphVisualization = ({ paths }) => {
    let nodes = [];
    let edges = [];

    if (!paths || paths.length === 1) {
        return <div>No path available!</div>;
    }

    nodes = [...new Set(paths.flat())].map((element, index) => ({
        id: `node_${index}`, 
        label: element,
        url: `${element}`
    }));

    for (let i = 0; i < paths.length - 1; i++) {
        const fromNode = nodes.find(node => node.label === paths[i].toString());
        const toNode = nodes.find(node => node.label === paths[i + 1].toString());

        if (fromNode && toNode) {
            edges.push({
                from: fromNode.id,
                to: toNode.id,
            });
        } else {
            console.warn(`Node not found for edge from ${paths[i]} to ${paths[i + 1]}`);
        }
    }

    const handleNodeClick = (event) => {
        const nodeId = event.nodes[0];
        const clickedNode = nodes.find(node => node.id === nodeId);
        if (clickedNode && clickedNode.url) {
            window.open(clickedNode.url, '_blank'); // Buka URL di tab baru
        }
    };

    const graph = {
        nodes,
        edges
    };

    const options = {
        layout: {
            hierarchical: false
        },
        nodes: {
            color: "#dcc9f1",
            font: {
                size: 14, 
                color: "#333"
            },
            shape: "box", 
            borderWidth: 2,
            borderWidthSelected: 4, 
            chosen: {
                node: function (values, id, selected, hovering) {
                    values.borderWidth = selected ? 4 : 2; 
                    values.borderColor = selected ? "#3a025e;" : "#dcc9f1";
                }
            }
        },
        edges: {
            color: "#611097",
            arrows: {
                to: {
                    enabled: true,
                    scaleFactor: 0.4
                }
            },
            width: 2, 
            hoverWidth: 0.5,
            smooth: {
                type: "continuous"
            }
        },
        interaction: {
            hover: true,
            navigationButtons: true,
            keyboard: true,
            selectConnectedEdges: false,
            hoverConnectedEdges: false,
            multiselect: false,
            tooltipDelay: 300,
            zoomView: true
        }
    };    

    return (
        <div style={{ height: '500px', width: '500px' }}>
            <Graph graph={graph} options={options} events={{ click: handleNodeClick }} />
        </div>
    );
};

export default GraphVisualization;
