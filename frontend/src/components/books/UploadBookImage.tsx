"use client";

import { Button } from "@/components/ui/button";
import { Upload, X } from "lucide-react";
import Image from "next/image";
import { useRef } from "react";

export default function UploadBookImage({
  bookImagePreview,
  setbookImagePreview,
  bookImageInfo,
  setBookImageInfo,
  bookImageError,
  setbookImageError,
  setImageFile,
  imageURL
}: {
  bookImagePreview: string | ArrayBuffer | null;
  setbookImagePreview: (file: string | ArrayBuffer | null) => void;
  bookImageInfo: { name: string; size: number } | null;
  setBookImageInfo: (info: { name: string; size: number } | null) => void;
  bookImageError: string | null;
  setbookImageError: (error: string | null) => void;
  setImageFile: (file: File | null) => void;
  imageURL?: string | null;
}) {

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      if (file.size > 4 * 1024 * 1024) {
        setbookImageError('File is too large. Please upload a file smaller than 4MB.');
        setbookImagePreview(null);
        setBookImageInfo(null);
        setImageFile(null); // Clear the image file
      } else {
        const reader = new FileReader();
        reader.onload = () => {
          setbookImagePreview(reader.result);
          setBookImageInfo({ name: file.name, size: file.size });
          setbookImageError(null);
          setImageFile(file); // Set the image file
        };
        reader.readAsDataURL(file);
      }
    }
  };

  const handleRemoveFile = (e: React.MouseEvent) => {
    setbookImagePreview(null);
    setBookImageInfo(null);
    setbookImageError(null);
    setImageFile(null); // Reset the image file
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
      <label className="form-label" htmlFor="image">Book Cover Image</label>
      <div className="h-[150px] w-full hover:cursor-pointer hover:bg-secondary/5 border-2 border-secondary border-dashed rounded-sm p-4" onClick={handleClick}>
        <input
          name="image"
          type="file"
          ref={fileInputRef}
          className="hidden"
          onChange={handleFileChange}
          accept="image/jpeg, image/png"
        />
        <div className="flex flex-col items-center justify-center w-full">
          {imageURL || bookImagePreview ? (
            <div className="flex flex-col items-center justify-center space-y-2">
              <div className="relative">
                <Image
                  width={56}
                  height={56}
                  src={bookImagePreview ? bookImagePreview.toString() : (imageURL || '/assets/books/book-placeholder.png')}
                  alt="preview" className="w-20 h-20 rounded-md shadow-md object-contain border" />
                <Button variant="link" className="absolute -top-3 -right-8 text-destructive" onClick={handleRemoveFile}>
                  <X className="h-6 w-6" />
                </Button>
              </div>
              <div className="flex flex-col items-center">
                <p className="text-xs text-primary font-medium">
                  {bookImageInfo && bookImageInfo?.name.length > 20
                    ? bookImageInfo?.name.slice(0, 20) + "..."
                    : bookImageInfo?.name}
                </p>
                <p className="text-xs text-primary font-medium">
                  {bookImageInfo?.size ? `${(bookImageInfo.size / 1024 / 1024).toFixed(2)} MB` : ""}
                </p>
              </div>
            </div>
          ) : (
            <div className="flex flex-col items-center justify-center space-y-2">
              <Upload className="h-7 w-7 text-primary" />
              <div className="flex flex-col items-center justify-center text-secondary space-y-2">
                <p className="text-sm underline underline-offset-4 text-primary">Upload Logo</p>
                <p className="text-xs text-primary/50">JPG, JPEG, PNG, less than 4MB</p>
              </div>
            </div>
          )}
        </div>
      </div>
      {bookImageError && <p className="text-sm text-destructive">{bookImageError}</p>}
    </div>
  );
}
