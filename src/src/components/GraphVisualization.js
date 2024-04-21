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
        label: element.toString(),
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


    const graph = {
        nodes,
        edges
    };

    const options = {
        layout: {
            hierarchical: false
        },
        edges: {
            color: "#000000",
            arrows: {
                to: {
                    enabled: true,
                    scaleFactor: 1 
                }
            }
        }
    };

    return (
        <div style={{ height: '500px' }}>
            <Graph graph={graph} options={options} />
        </div>
    );
};

export default GraphVisualization;
