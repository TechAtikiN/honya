"use client";

import { Button } from "@/components/ui/button";
import { Upload, X } from "lucide-react";
import Image from "next/image";
import { useRef, useState } from "react";

export default function UploadBookImage() {
  const [filePreview, setFilePreview] = useState<string | ArrayBuffer | null>(null);
  const [fileInfo, setFileInfo] = useState<{ name: string; size: number } | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      if (file.size > 4 * 1024 * 1024) {
        setError("File is too large. Please upload a file smaller than 4MB.");
        setFilePreview(null);
        setFileInfo(null);
      } else {
        const reader = new FileReader();
        reader.onload = () => {
          setFilePreview(reader.result);
          setFileInfo({ name: file.name, size: file.size });
          setError(null); // Clear error if the file is valid
        };
        reader.readAsDataURL(file);
      }
    }
  };

  const handleRemoveFile = (e: React.MouseEvent) => {
    setFilePreview(null);
    setFileInfo(null);
    setError(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = ""; // Reset the file input value so that the user can upload the file again
    }
    e.stopPropagation();
    e.preventDefault();
  };

  const fileInputRef = useRef<HTMLInputElement | null>(null);

  const handleClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  return (
    <div className="flex flex-col space-y-2 col-span-2 ">
      <label className="form-label" htmlFor="logo">
        Book Cover Image
      </label>
      <div
        className="h-[150px] w-full hover:cursor-pointer hover:bg-secondary/5 border-2 border-secondary border-dashed rounded-sm p-4"
        onClick={handleClick}
      >
        <input
          type="file"
          ref={fileInputRef}
          className="hidden"
          onChange={handleFileChange}
          accept="image/jpeg, image/png"
        />
        <div className="flex flex-col items-center justify-center w-full">
          {filePreview ? (
            <div className="flex flex-col items-center justify-center space-y-2">
              <div className="relative">
                <Image
                  width={56}
                  height={56}
                  src={filePreview as string}
                  alt="preview"
                  className="w-14 h-14 rounded-full shadow-md object-cover border"
                />
                <Button
                  variant="link"
                  className="absolute -top-3 -right-8 text-destructive"
                  onClick={handleRemoveFile}
                >
                  <X className="h-6 w-6" />
                </Button>
              </div>
              <div className="flex flex-col items-center">
                <p className="text-xs text-primary font-medium">
                  {fileInfo && fileInfo?.name.length > 20
                    ? fileInfo?.name.slice(0, 20) + "..."
                    : fileInfo?.name}
                </p>
                <p className="text-xs text-primary font-medium">
                  {fileInfo?.size ? `${(fileInfo.size / 1024 / 1024).toFixed(2)} MB` : ""}
                </p>
              </div>
            </div>
          ) : (
            <div className="flex flex-col items-center justify-center space-y-2">
              <Upload className="h-7 w-7 text-primary" />
              <div className="flex flex-col items-center justify-center text-secondary space-y-2">
                <p className="text-sm underline underline-offset-4 text-primary">
                  Upload Logo
                </p>
                <p className="text-xs text-primary/50">JPG, JPEG, PNG, less than 4MB</p>
              </div>
            </div>
          )}
        </div>
      </div>
      {error && (
        <p className="text-sm text-destructive">{error}</p>
      )}
    </div>
  );
}
