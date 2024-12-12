export const STYLE_MAP: Record<string, { shape: string; color: string }> = {
    start: { shape: "ellipse", color: "#3343ff" }, // Start node as ellipse (oval)
    end: { shape: "ellipse", color: "#ff3333" }, // End node as ellipse (oval)
    decision: { shape: "diamond", color: "#5d33ff" }, // Decision node as diamond
    process: { shape: "box", color: "#c333ff" }, // Process node as rectangle (box)
    input: { shape: "parallelogram", color: "#33c3ff" }, // Input node as parallelogram
    output: { shape: "parallelogram", color: "#33c3ff" }, // Output node as parallelogram
    preparation: { shape: "rect", color: "#ffcc33" }, // Preparation node as rectangle
    connector: { shape: "ellipse", color: "#ff99cc" }, // Connector node as ellipse (link)
    manual: { shape: "rect", color: "#ccff33" }, // Manual node as rectangle
    storage: { shape: "hexagon", color: "#ff9933" }, // Storage node as hexagon
  };
