"use client";

import React, { useEffect, useRef, useMemo } from "react";
import { Network, Edge, Node } from "vis-network";
import { DataSet } from "vis-data";

export type FlowChartNode = Node & {
  type?: "start" | "end" | "decision" | "process";
};

export type FlowChartData = {
  nodes: FlowChartNode[];
  edges: Edge[];
};

type FlowChartProps = {
  data: FlowChartData;
};

const FlowChart: React.FC<FlowChartProps> = ({ data }) => {
  const containerRef = useRef<HTMLDivElement>(null);

  const memoizedData = useMemo(() => {
    // Map node types to specific shapes
    const shapeMap: Record<string, string> = {
      start: "box",
      end: "box",
      decision: "diamond",
      process: "ellipse",
    };

    // Assign shapes based on type
    const updatedNodes = data.nodes.map((node) => ({
      ...node,
      shape: shapeMap[node.type || "process"], // Default to "process" shape
    }));

    return { nodes: updatedNodes, edges: data.edges };
  }, [data]);

  useEffect(() => {
    if (containerRef.current) {
      const nodes = new DataSet(memoizedData.nodes);
      const edges = new DataSet(memoizedData.edges);

      const network = new Network(
        containerRef.current,
        { nodes, edges },
        {
          layout: { hierarchical: false },
          interaction: { dragNodes: true },
          physics: { enabled: true },
          edges: {
            font: { align: "top" },
            arrows: { to: { enabled: true, scaleFactor: 1 } },
          },
        }
      );

      return () => {
        network.destroy(); // Clean up on unmount
      };
    }
  }, [memoizedData]);

  return (
    <div
      ref={containerRef}
      className="h-[500px] w-full overflow-auto border-2 border-gray-300 rounded-lg shadow-md"
    />
  );
};

export default FlowChart;
