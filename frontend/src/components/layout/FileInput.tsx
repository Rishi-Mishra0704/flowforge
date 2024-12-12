import React from "react";

interface FileInputProps {
  className?: string;
  onFileChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const FileInput: React.FC<FileInputProps> = ({ onFileChange, className }) => {
  return (
    <div>
      <input
        onChange={onFileChange}
        type="file"
        accept=".zip"
        className={`file-input file-input-bordered file-input-accent w-full max-w-xs} ${className}`}
      />
    </div>
  );
};

export default FileInput;
