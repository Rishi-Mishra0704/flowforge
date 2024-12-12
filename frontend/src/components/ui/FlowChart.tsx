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
  if (typeof window === "undefined") {
    return null;
  }

  if (!data || data === undefined) {
    return null;
  }

  const containerRef = useRef<HTMLDivElement>(null);

  const memoizedData = useMemo(() => {
    // Map node types to specific shapes and colors
    const styleMap: Record<string, { shape: string; color: string }> = {
      start: { shape: "box", color: "#3343ff" },
      end: { shape: "box", color: "#ff3333" },
      decision: { shape: "diamond", color: "#5d33ff" },
      process: { shape: "ellipse", color: "#c333ff" },
    };

    // Transform nodes and apply styles
    const updatedNodes = data.nodes.map((node) => {
      if (!node.type) {
        throw new Error(`Node type is missing for node: ${node.id}`);
      }
      const typeKey = node.type.toLowerCase();
      return {
        id: node.id,
        label: node.label,
        shape: styleMap[typeKey]?.shape,
        color: { background: styleMap[typeKey]?.color || "#c333ff" }, // Default color
        font: { color: "#fff" }, // White text for contrast
      };
    });

    // Transform edges to match `vis-network` expectations
    const updatedEdges = data.edges.map((edge, index) => ({
      id: index,
      from: edge.from,
      to: edge.to,
    }));

    return { nodes: updatedNodes, edges: updatedEdges };
  }, [data]);

  useEffect(() => {
    if (containerRef.current) {
      const nodes = new DataSet(memoizedData.nodes);
      const edges = new DataSet(memoizedData.edges);

      const network = new Network(
        containerRef.current,
        { nodes, edges },
        {
          layout: { hierarchical: true },
          interaction: { dragNodes: true, dragView: true, hover: true },
          physics: { enabled: false }, // Disable physics for static layout
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
