import React from "react";
import FileInput from "@/components/ui/FileInput";
import FlowChart, { FlowChartData } from "@/components/ui/FlowChart";

const data: FlowChartData = {
  nodes: [
    { id: 1, label: "Start", type: "start" },
    { id: 2, label: "Decision", type: "decision" },
    { id: 3, label: "Process", type: "process" },
    { id: 4, label: "End", type: "end" },
  ],
  edges: [
    { from: 1, to: 2 },
    { from: 2, to: 3 },
    { from: 3, to: 4 },
  ],
};

const FlowchartPage: React.FC = () => {
  return (
    <div className="p-4 space-y-4">
      <FileInput />
      <div className="p-4 bg-white border-2 border-gray-300 rounded-lg shadow-lg">
        <h2 className="text-xl font-semibold text-gray-700 mb-4">
          Interactive Flowchart
        </h2>
        <FlowChart data={data} />
      </div>
    </div>
  );
};

export default FlowchartPage;
