import { Button } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { BsCamera, BsCart, BsTrash } from "react-icons/bs";
import { HiOutlineDocumentAdd } from "react-icons/hi";
import { IoDocumentsOutline } from "react-icons/io5";
import { FileModel } from "../models/file";
import { ProductModel } from "../models/product";
import { uploadFile } from "../services/api/commonApi";
import { money } from "../utils/helper";
import MessageMention from "./MessageMention";
import { parseMentions } from "../utils/helper-ui";
import { CiBoxList } from "react-icons/ci";
import {
  WhatsappInteractiveListRow,
  WhatsappInteractiveListSection,
  WhatsappInteractiveModel,
} from "../models/whatsapp_interactive_message";
import { getInteractiveTemplate } from "../services/api/templateApi";
import toast from "react-hot-toast";
import { Link } from "react-router-dom";
import { AiOutlineFile, AiOutlineLink } from "react-icons/ai";

interface MessageTemplateFieldProps {
  index: number;
  title: string;
  body: string;
  onChangeBody: (val: string) => void;
  onClickEmoji: () => void;
  files: FileModel[];
  product?: ProductModel;
  interactive?: WhatsappInteractiveModel;
  onUploadFile: (file: FileModel, index?: number) => void;
  onUploadImage: (file: FileModel, index?: number) => void;
  showDelete?: boolean;
  onDelete?: () => void;
  onTapProduct?: () => void;
  onTapInteractive?: () => void;
  onEditInteractive?: (d: WhatsappInteractiveModel) => void;
  readonly?: boolean;
  onDeleteFile?: (file: FileModel) => void;
  onDeleteImage?: (file: FileModel) => void;
  disableProduct?: boolean;
  disableInteractive?: boolean;
  msgId?: string;
  templateId?: string;
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
  disableInteractive,
  onTapInteractive,
  onEditInteractive,
  interactive,
  msgId,
  templateId,
}) => {
  const [selectedInteractive, setSelectedInteractive] =
    useState<WhatsappInteractiveModel>();
  useEffect(() => {
    getInteractiveTemplate(templateId!, msgId)
      .then((res: any) => {
        console.log(res);
        setSelectedInteractive(res.data);
      })
      .catch(toast.error);
  }, [msgId]);
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
                  <span>
                    {" "}
                    {readonly ? "No Photo" : "Add Photo to message"}{" "}
                  </span>
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
      {!disableInteractive && (
        <div className="mt-8">
          <h4 className="font-semibold">Interactive</h4>
          <div className="flex flex-col justify-center items-center p-4 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100">
            {!selectedInteractive ? (
              <div
                className="flex flex-col items-center text-center cursor-pointer"
                onClick={() => {
                  if (readonly) return;
                  onTapInteractive?.();
                }}
              >
                <span>
                  {" "}
                  {readonly
                    ? "No Interactive Message"
                    : "Add Interactive Message"}{" "}
                </span>
                {readonly ? null : <CiBoxList size={32} />}
              </div>
            ) : (
              <div
                className="w-full"
                onClick={() => {
                  onEditInteractive?.(selectedInteractive);
                }}
              >
                <table className="w-full">
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Title</td>
                    <td>{selectedInteractive?.title}</td>
                  </tr>
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Description</td>
                    <td>{selectedInteractive?.description}</td>
                  </tr>
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Type</td>
                    <td>{selectedInteractive?.type}</td>
                  </tr>
                  <tr className="border-b">
                    <td className="font-semibold w-1/4 py-2">Data</td>
                    <td></td>
                  </tr>
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Header</td>
                    <td>
                      {selectedInteractive?.data?.header?.type == "image" && (
                        <img
                          src={selectedInteractive?.data?.header?.image?.link}
                          alt="header"
                          className="w-32 h-32 rounded-lg object-cover"
                        />
                      )}
                      {selectedInteractive?.data?.header?.type == "video" && (
                        <video
                          src={selectedInteractive?.data?.header?.video?.link}
                          controls
                        />
                      )}
                      {selectedInteractive?.data?.header?.type ==
                        "document" && (
                        <Link
                          to={selectedInteractive?.data?.header?.document?.link}
                        >
                          <AiOutlineFile size={32} />
                        </Link>
                      )}
                      {selectedInteractive?.data?.header?.type == "text" && (
                        <span>{selectedInteractive?.data?.header?.text}</span>
                      )}
                    </td>
                  </tr>
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Body</td>
                    <td>{selectedInteractive?.data?.body?.text}</td>
                  </tr>
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Footer</td>
                    <td>{selectedInteractive?.data?.footer?.text}</td>
                  </tr>
                  {selectedInteractive?.data?.type == "list" && (
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Action</td>
                    <td>{selectedInteractive?.data?.action?.button}</td>
                  </tr>
                  )}
                  {selectedInteractive?.data?.type == "cta_url" && (
                  <tr>
                    <td className="font-semibold w-1/4 py-2">Action</td>
                    <td className="flex">{selectedInteractive?.data?.action?.parameters?.display_text}  <AiOutlineLink size={16} onClick={() => window.open(selectedInteractive?.data?.action?.parameters?.url, "_blank")}/></td>
                  </tr>
                  )}
                  {selectedInteractive?.data?.type == "list" && (
                    <tr className="border-b">
                      <td className="font-semibold w-1/4 py-2">Sections</td>
                      <td></td>
                    </tr>
                  )}
                </table>
                {selectedInteractive?.data?.type == "list" &&
                  selectedInteractive?.data?.action?.sections?.map(
                    (
                      section: WhatsappInteractiveListSection,
                      index: number
                    ) => (
                      <div key={index} className="w-full mt-4">
                        <h4 className="font-semibold">{section.title}</h4>
                        <table className="w-full ">
                          <thead>
                            <tr>
                              <th className="p-2 border border-gray-100">
                                Title
                              </th>
                              <th className="p-2 border border-gray-100">
                                Description
                              </th>
                            </tr>
                          </thead>
                          <tbody>
                            {section.rows.map(
                              (row: WhatsappInteractiveListRow) => (
                                <tr key={row.id}>
                                  <td className="p-2 border border-gray-100">
                                    {row.title}
                                  </td>
                                  <td className="p-2 border border-gray-100">
                                    {row.description}
                                  </td>
                                </tr>
                              )
                            )}
                          </tbody>
                        </table>
                      </div>
                    )
                  )}
              </div>
            )}
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
