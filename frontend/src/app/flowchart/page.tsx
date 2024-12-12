"use client";
import React, { useState } from "react";
import axios from "axios";
import FlowChart, { FlowChartData } from "@/components/layout/FlowChart";
import { Alert } from "@/components/ui/alert";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import FileInput from "@/components/layout/FileInput";

const FlowchartPage: React.FC = () => {
  const [flowchartData, setFlowchartData] = useState<FlowChartData | null>(
    null
  );
  const [file, setFiles] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleFileUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const uploadedFile = event.target.files?.[0];
    setFiles(uploadedFile || null);
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
      setFiles(null);
    } catch (err: any) {
      setError(
        err.response?.data?.message || "Failed to upload and process file"
      );
      setFiles(null);
    }
  };

  return (
    <div className="p-2 space-y-6 bg-gray-100 dark:bg-gray-900 text-gray-800 dark:text-gray-100 min-h-screen">
      <Card className="border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800">
        <CardHeader>
          <CardTitle className="text-lg font-medium text-gray-700 dark:text-gray-200">
            Upload Codebase to Generate Flowchart
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex flex-col sm:flex-row sm:items-center space-y-4 sm:space-y-0 sm:space-x-4">
            <FileInput
              onFileChange={handleFileUpload}
              className="btn-primary dark:bg-primary dark:text-gray-100"
            />
            {error && (
              <Alert
                variant="destructive"
                className="text-red-500 dark:text-red-400"
              >
                {error}
              </Alert>
            )}
          </div>
        </CardContent>
      </Card>

      {flowchartData && (
        <Card className="border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800">
          <CardHeader>
            <CardTitle className="text-gray-700 dark:text-gray-200">
              Interactive Flowchart
            </CardTitle>
          </CardHeader>
          <CardContent>
            <FlowChart data={flowchartData} />
          </CardContent>
        </Card>
      )}
    </div>
  );
};

export default FlowchartPage;
