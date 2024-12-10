import React from "react";

interface FileInputProps {
  onFileChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const FileInput: React.FC<FileInputProps> = ({onFileChange}) => {
  return (
    <div>
      <input
        onChange={onFileChange}
        type="file"
        className="file-input file-input-bordered file-input-accent w-full max-w-xs"
      />
    </div>
  );
};

export default FileInput;
