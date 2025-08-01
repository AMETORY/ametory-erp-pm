import { Button } from "flowbite-react";
import type { FC } from "react";
import { BsCamera, BsCart, BsTrash } from "react-icons/bs";
import { HiOutlineDocumentAdd } from "react-icons/hi";
import { IoDocumentsOutline } from "react-icons/io5";
import { FileModel } from "../models/file";
import { ProductModel } from "../models/product";
import { uploadFile } from "../services/api/commonApi";
import { money } from "../utils/helper";
import MessageMention from "./MessageMention";
import { parseMentions } from "../utils/helper-ui";

interface MessageTemplateFieldProps {
  index: number;
  title: string;
  body: string;
  onChangeBody: (val: string) => void;
  onClickEmoji: () => void;
  files: FileModel[];
  product?: ProductModel;
  onUploadFile: (file: FileModel, index?: number) => void;
  onUploadImage: (file: FileModel, index?: number) => void;
  showDelete?: boolean;
  onDelete?: () => void;
  onTapProduct?: () => void;
  readonly?: boolean;
  onDeleteFile?: (file: FileModel) => void;
  onDeleteImage?: (file: FileModel) => void;
  disableProduct?: boolean;
}

const MessageTemplateField: FC<MessageTemplateFieldProps> = ({
  index,
  title,
  body,
  onChangeBody,
  onClickEmoji,
  onUploadImage,
  onUploadFile,
  files,
  product,
  onDelete,
  onTapProduct,
  readonly,
  onDeleteFile,
  onDeleteImage,
  disableProduct,
}) => {
  return (
    <div className="bg-gray-50 rounded-lg p-4 flex flex-col mb-8">
      <h4 className="font-semibold">{title}</h4>
      {readonly ? (
        <div className="p-4 bg-white">
          {parseMentions(body ?? "", (type, id) => {})}
        </div>
      ) : (
        <MessageMention
          msg={body}
          onChange={(val: any) => {
            onChangeBody(val.target.value);
          }}
          onClickEmoji={onClickEmoji}
          onSelectEmoji={(emoji: string) => {
            onChangeBody(`${body}${emoji}`);
          }}
        />
      )}

      <div className="mt-8">
        <h4 className="font-semibold">Files</h4>
        <div className="grid grid-cols-2 gap-4">
          <div className="flex flex-col justify-center items-center p-2 rounded-lg bg-white relative">
            <div
              className="cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100 w-full h-full flex justify-center items-center p-16"
              onClick={() => {
                if (readonly) return;
                document.getElementById(`image-${index}`)?.click();
              }}
            >
              {files.filter((f) => f.mime_type.includes("image")).length ===
              0 ? (
                <div className="flex flex-col items-center text-center">
                  <span> {readonly ? "No Photo" : "Add Photo to message"} </span>
                  {readonly ? null : <BsCamera />}
                </div>
              ) : (
                <img
                  className="w-32 h-32 object-cover"
                  src={files.find((f) => f.mime_type.includes("image"))?.url}
                />
              )}
            </div>

            <input
              type="file"
              className="hidden"
              accept="image/*"
              id={`image-${index}`}
              onChange={(e) => {
                const file = e.target.files?.[0];
                if (file) {
                  uploadFile(file, {}, () => {}).then((resp: any) => {
                    onUploadImage(resp.data, index);
                  });
                }
              }}
            />
            {files.find((f) => f.mime_type.includes("image")) && (
              <BsTrash
                size={20}
                className="absolute bottom-2 right-2 cursor-pointer text-red-400 hover:text-red-600"
                onClick={() => {
                  onDeleteImage?.(
                    files.find((f) => f.mime_type.includes("image"))!
                  );
                }}
              />
            )}
          </div>
          <div className="flex flex-col justify-center items-center p-2 rounded-lg bg-white relative">
            <div
              className="cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100 w-full h-full flex justify-center items-center p-16"
              onClick={() => {
                if (readonly) return;
                document.getElementById(`image-${index}-file`)?.click();
              }}
            >
              {files.filter((f) => !f.mime_type.includes("image")).length ===
              0 ? (
                <div className="flex flex-col items-center text-center">
                  <span>{readonly ? "No File" : "Add File to message"} </span>
                  {readonly ? null : <HiOutlineDocumentAdd size={32} />}
                </div>
              ) : (
                // <IoAttach className="rotate-[30deg]" size={32}/>
                <div className="flex items-center flex-col px-8">
                  <IoDocumentsOutline size={32} />
                  <small className="text-center mt-4">
                    {
                      files.find((f) => !f.mime_type.includes("image"))
                        ?.file_name
                    }
                  </small>
                </div>
              )}
            </div>
            <input
              type="file"
              className="hidden"
              accept=".doc,.docx,.pdf,.xls,.xlsx,.txt"
              id={`image-${index}-file`}
              onChange={(e) => {
                const file = e.target.files?.[0];
                if (file) {
                  uploadFile(file, {}, () => {}).then((resp: any) => {
                    onUploadFile(resp.data, index);
                  });
                }
              }}
            />
            {files.find((f) => !f.mime_type.includes("image")) && (
              <BsTrash
                size={20}
                className="absolute bottom-2 right-2 cursor-pointer text-red-400 hover:text-red-600"
                onClick={() => {
                  onDeleteFile?.(
                    files.find((f) => !f.mime_type.includes("image"))!
                  );
                }}
              />
            )}
          </div>
        </div>
      </div>
      {!disableProduct && (
        <div className="mt-8">
          <h4 className="font-semibold">Product</h4>
          <div className="grid grid-cols-2 gap-4 ">
            <div
              className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100"
              onClick={() => {
                if (readonly) return;
                onTapProduct?.();
              }}
            >
              {!product ? (
                <div className="flex flex-col items-center text-center">
                  <span> {readonly ? "No Product" : "Add Product"} </span>
                  {readonly ? null : <BsCart size={32} />}
                </div>
              ) : (
                <div className="flex items-center flex-col  px-8">
                  {(product?.product_images ?? []).length > 0 && (
                    <img
                      src={product?.product_images![0].url}
                      alt="product"
                      className="w-32 h-32 rounded-lg"
                    />
                  )}
                  <h3 className="font-semibold mt-2 text-center">
                    {product?.name}
                  </h3>
                  <small>{money(product?.price)}</small>
                </div>
              )}
            </div>
          </div>
        </div>
      )}
      <div className="mt-4">
        {onDelete && (
          <Button
            className=""
            color="red"
            onClick={() => {
              onDelete();
            }}
          >
            + Delete Message
          </Button>
        )}
      </div>
    </div>
  );
};
export default MessageTemplateField;
