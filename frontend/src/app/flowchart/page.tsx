"use client";
import React, { useState } from "react";
import axios from "axios";
import FlowChart, { FlowChartData } from "@/components/ui/FlowChart";
import FileInput from "@/components/ui/FileInput";

const FlowchartPage: React.FC = () => {
  const [flowchartData, setFlowchartData] = useState<FlowChartData | null>(
    null
  );
  const [error, setError] = useState<string | null>(null);

  const handleFileUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
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
      setError(err.response?.data?.message || "Failed to upload and process file");
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
