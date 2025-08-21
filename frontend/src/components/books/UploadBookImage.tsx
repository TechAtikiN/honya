"use client";

import { Button } from "@/components/ui/button";
import { Upload, X } from "lucide-react";
import { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";

export default function UploadBookImage() {
  const [filePreview, setFilePreview] = useState<string | ArrayBuffer | null>(
    null
  );

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = new FileReader();
    file.onload = function () {
      setFilePreview(file.result);
    };

    file.readAsDataURL(acceptedFiles[0]);
  }, []);

  const { acceptedFiles, getRootProps, getInputProps } = useDropzone({
    onDrop,
  });

  return (
    <div className="flex flex-col space-y-2 col-span-2 ">
      <label className="form-label" htmlFor="logo">
        Book Cover Image
      </label>
      <div
        className="h-[150px] w-full hover:cursor-pointer hover:bg-secondary/5 border-2 border-secondary border-dashed rounded-sm p-4"
        {...getRootProps()}
      >
        <input className="" type="file" {...getInputProps()} />

        <div className="flex flex-col items-center justify-center w-full">
          {filePreview ? (
            <div className="flex flex-col items-center justify-center space-y-2">
              <div className="relative">
                <img
                  src={filePreview as string}
                  alt="preview"
                  className="w-14 h-14 rounded-full shadow-md object-cover border"
                />
                <Button
                  variant="link"
                  className="absolute -top-3 -right-8 text-destructive"
                  onClick={(e) => {
                    setFilePreview(null);
                    e.stopPropagation();
                    e.preventDefault();
                  }}
                >
                  <X className="h-6 w-6" />
                </Button>
              </div>
              <p className="text-sm text-secondary">
                {acceptedFiles[0]?.name.length > 20
                  ? acceptedFiles[0]?.name.slice(0, 20) + "..."
                  : acceptedFiles[0]?.name}
              </p>
            </div>
          ) : (
            <div className="flex flex-col items-center justify-center space-y-2">
              <Upload className="h-7 w-7 text-primary" />
              <div className="flex flex-col items-center justify-center text-secondary space-y-2">
                <p className="text-sm underline underline-offset-4 text-primary">
                  Upload Logo
                </p>
                <p className="text-xs text-primary/50">JPG, JPEG, PNG, less than 5MB</p>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
