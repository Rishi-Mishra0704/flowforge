"use client";
import React, { useState } from "react";
import axios from "axios";
import FlowChart, { FlowChartData } from "@/components/ui/FlowChart";
import FileInput from "@/components/ui/FileInput";
const data: FlowChartData = {
  nodes: [
    { id: 1, label: "Start", type: "start" },
    { id: 2, label: "Is Student", type: "decision" },
    { id: 3, label: "Process A", type: "process" },
    { id: 4, label: "Process B", type: "process" },
    { id: 5, label: "End A", type: "end" },
    { id: 6, label: "End B", type: "end" },
    { id: 7, label: "Decision B", type: "decision" },
    { id: 8, label: "Process C", type: "process" },
    { id: 9, label: "End C", type: "end" },
  ],
  edges: [
    { from: 1, to: 2 },
    { from: 2, to: 3, label: "Yes" },
    { from: 2, to: 4, label: "No" },
    { from: 3, to: 5 },
    { from: 4, to: 7 },
    { from: 7, to: 6, label: "Yes" },
    { from: 7, to: 8, label: "No" },
    { from: 8, to: 9 },
  ],
};

const FlowchartPage: React.FC = () => {
  const [flowchartData, setFlowchartData] = useState<FlowChartData | null>(
    null
  );
  const [error, setError] = useState<string | null>(null);

  const handleFileUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("codebase", file);

    try {
      setError(null); // Clear previous errors
      const response = await axios.post(
        "http://localhost:8080/flowchart", // Replace with your backend URL
        formData,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      );
      setFlowchartData(response.data);
    } catch (err: any) {
      setError(
        err.response?.data?.message || "Failed to upload and process file"
      );
    }
  };

  return (
    <div className="p-4 space-y-4">
      <div>
        <FileInput onFileChange={handleFileUpload} />
        {error && <p className="text-red-500 mt-2">{error}</p>}
      </div>
      {flowchartData && (
      <div className="p-4 bg-white border-2 border-gray-300 rounded-lg shadow-lg">
        <h2 className="text-xl font-semibold text-gray-700 mb-4">
          Interactive Flowchart
        </h2>
        <FlowChart data={flowchartData} />
      </div>
      )}
    </div>
  );
};

export default FlowchartPage;
