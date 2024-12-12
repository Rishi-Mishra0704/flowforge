import { FlowChartData } from "@/components/layout/FlowChart";

export const data: FlowChartData = {
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