import React from 'react';
import Graph from 'react-graph-vis';

const GraphVisualization = ({paths, updateTrigger}) => {
    let nodes = [];
    let edges = [];

    if (paths[0] === null) {
        return <div>No path available!</div>;
    }

    nodes = [...new Set(paths.flat())].map((element, index) => ({
        id: `node_${index}`, 
        label: element.split('/').pop().replace(/_/g, " "),
        url: `${element}`
    }));    

    for (let i = 0; i < paths.length; i++) {
        for (let j = 0; j < paths[i].length - 1; j++) {
            const fromNode = nodes.find(node => node.url === paths[i][j]);
            const toNode = nodes.find(node => node.url === paths[i][j + 1]);

            if (fromNode && toNode) {
                edges.push({
                    from: fromNode.id,
                    to: toNode.id,
                });
            } else {
                console.warn(`Node not found for edge from ${paths[i][j]} to ${paths[i][j + 1]}`);
            }
        }
    }

    const handleNodeClick = (event) => {
        const nodeId = event.nodes[0];
        const clickedNode = nodes.find(node => node.id === nodeId);
        if (clickedNode && clickedNode.url) {
            window.open(clickedNode.url, '_blank');
        }
    };

    const graph = {
        nodes,
        edges
    };

    const options = {
        layout: {
            hierarchical: {
                direction: "UD",
                sortMethod: "directed",
                levelSeparation: 100,
            },
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
            zoomView: true,
            initialZoomLevel: 2,
            maxZoomLevel: 3
        }
    };    

    return (
        <div style={{ height: '500px', width: '1000px' }}>
            <Graph graph={graph} options={options} events={{ click: handleNodeClick }} />
        </div>
    );
};

export default GraphVisualization;
