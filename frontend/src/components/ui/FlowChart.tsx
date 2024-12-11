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
    // Map node types to specific shapes and colors
    const styleMap: Record<string, { shape: string; color: string }> = {
      start: { shape: "box", color: "#3343ff" }, // Soft Mint Green for Start
      end: { shape: "box", color: "#ff3333" }, // Soft Rose for End
      decision: { shape: "circle", color: "#5d33ff" }, // Pastel Orange for Decision
      process: { shape: "ellipse", color: "#c333ff" }, // Light Sky Blue for Process
    };
    

    // Assign shapes and colors based on type
    const updatedNodes = data.nodes.map((node) => ({
      ...node,
      shape: styleMap[node.type || "process"].shape, // Default to "process" shape
      color: { background: styleMap[node.type || "process"].color }, // Default color
      font: { color: "#fff" }, // White text for better contrast
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
          layout: { hierarchical: { direction: "UD", sortMethod: "directed" } }, // Top to bottom
          interaction: { dragNodes: true, dragView: true, hover: true },
          physics: { enabled: false }, // Disable physics for more static layout
          edges: {
            font: { align: "middle" },
            arrows: { to: { enabled: true, scaleFactor: 1.2 } }, // Enhance arrow size
            color: { color: "#999999", highlight: "#000000" }, // Gray edges
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
