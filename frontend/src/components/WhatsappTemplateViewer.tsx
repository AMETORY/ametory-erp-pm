import { useRef, type FC } from "react";
import { WhatsappAPITemplate } from "../models/whatsapp_api_template";
import { BsCamera } from "react-icons/bs";
import Markdown from "react-markdown";
import { uploadFile } from "../services/api/commonApi";

interface WhatsappTemplateViewerProps {
  template: WhatsappAPITemplate;
  whatsappTemplateMappingParams?: any[];
  onChangeWhatsappTemplateMappingParams?: (params: any[]) => void;
  headerImageUrl?: string;
  onChangeHeaderImageUrl?: (url: string) => void;
  whatsappTemplateID?: string;
  onWhatsappTemplateChange?: (id: string) => void;
  isView?: boolean;
}

const WhatsappTemplateViewer: FC<WhatsappTemplateViewerProps> = ({
  template,
  whatsappTemplateMappingParams = [],
  onChangeWhatsappTemplateMappingParams,
  headerImageUrl,
  onChangeHeaderImageUrl,
  whatsappTemplateID,
  onWhatsappTemplateChange,
  isView,
}) => {
  const fileRef = useRef<HTMLInputElement>(null);
  const msgParams = [
    { id: "{{user}}", display: "Full Name" },
    { id: "{{phone}}", display: "Phone Number" },
    { id: "{{agent}}", display: "Agent Name" },
    { id: "{{product}}", display: "Product" },
  ];
  let header = template.components.find((c) => c.type == "HEADER");
  let body = template.components.find((c) => c.type == "BODY");
  let footer = template.components.find((c) => c.type == "FOOTER");
  let textHeader = null;
  let imageHeader = null;
  if (header && header.format == "IMAGE") {
    imageHeader = header!.example!.header_handle![0];
    if (headerImageUrl) {
      imageHeader = headerImageUrl;
    }
  }
  if (header && header.format == "TEXT") {
    textHeader = header!.text;
  }

  let params = [];

  let bodyText = null;
  if (body) {
    bodyText = body!.text;
    if (body?.example?.body_text) {
      for (let index = 0; index < body?.example?.body_text[0].length; index++) {
        const element = body?.example?.body_text[0][index];
        let placeHolder = "{{" + (index + 1) + "}}";
        bodyText = bodyText?.replace(placeHolder, element);
        params.push(placeHolder);
      }
    }

    if (body?.example?.body_text_named_params) {
      for (
        let index = 0;
        index < body?.example?.body_text_named_params.length;
        index++
      ) {
        const element = body?.example?.body_text_named_params[index];
        let placeHolder = "{{" + element.param_name + "}}";
        bodyText = bodyText?.replace(placeHolder, element.example);
        params.push(placeHolder);
      }
    }

    // console.log(body?.example?.body_text_named_params);
  }
  let footerText = null;
  if (footer) {
    footerText = footer!.text;
  }


  return (
    <div
      className={`border border-gray-200 bg-white rounded-lg cursor-pointer hover:border-gray-400 h-fit`}
      key={template.id}
      style={{
        borderColor:
          whatsappTemplateID == template.name ? "#3b82f6" : "#f3f3f3",
        borderWidth: whatsappTemplateID == template.name ? "3px" : "1px",
      }}
      onClick={() => {
        onWhatsappTemplateChange?.(template.name);
      }}
    >
      {imageHeader && (
        <div className="relative">
          <img
            src={imageHeader}
            alt={template.name}
            className="w-full rounded-t"
          />
          {!isView && (
            <div
              className="absolute bottom-2 left-2 bg-[rgba(0,0,0,0.5)] rounded-full p-2 flex gap-2 items-center"
              onClick={() => fileRef.current?.click()}
            >
              <BsCamera className="w-4 h-4 text-white" />
              <span className="text-sm text-white">Upload Image</span>
            </div>
          )}
        </div>
      )}

      <div className="p-4">
        {textHeader && (
          <h3 className="text-lg font-semibold mb-2 leading-tight">
            <Markdown>{textHeader}</Markdown>
          </h3>
        )}
        {bodyText && (
          <p className="text-gray-700 text-sm leading-snug">
            <Markdown>{bodyText}</Markdown>
          </p>
        )}

        {footerText && (
          <p className="text-gray-400 text-xs leading-snug">{footerText}</p>
        )}
      </div>
      {params.length > 0 && whatsappTemplateID == template.name && (
        <div className="bg-gray-100 p-4">
          <div className="overflow-x-auto">
            <table className="table-auto w-full text-sm">
              <thead>
                <tr className="border-b">
                  <th className="px-2 py-2">Parameter</th>
                  <th className="px-2 py-2">Mapping Param</th>
                </tr>
              </thead>
              <tbody>
                {params.map((param, index) => {
                  let value = whatsappTemplateMappingParams?.find(
                    (p) => p.param == param
                  );
                  return (
                    <tr key={index} className="border-b last:border-b-0">
                      <td className="px-2 py-2">{param}</td>
                      <td className="px-2 py-2">
                        <select
                          className="w-full px-2 py-1 border border-gray-300 rounded-md"
                          disabled={isView}
                          value={value ? value.value : ""}
                          onChange={(e) => {
                            let params = [...whatsappTemplateMappingParams];
                            // console.log("params", whatsappTemplateMappingParams);
                            // onParamChange?.(param, e.target.value);
                            let paramsIndex =
                              whatsappTemplateMappingParams?.findIndex(
                                (p) => p.param == param
                              );
                            // console.log("paramsIndex",paramsIndex);
                            if (paramsIndex < 0) {
                              params.push({
                                param: param,
                                value: e.target.value,
                              });
                            } else {
                              params[paramsIndex] = {
                                param: param,
                                value: e.target.value,
                              };
                            }

                            // console.log("params", params);
                            onChangeWhatsappTemplateMappingParams?.(params);
                          }}
                        >
                          <option value="">Select Param</option>
                          {msgParams.map((param) => (
                            <option key={param.id} value={param.id}>
                              {param.display}
                            </option>
                          ))}
                          
                        </select>
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
      )}
      <input
        accept=".png, .jpg, .jpeg"
        type="file"
        name="file"
        id=""
        ref={fileRef}
        className="hidden"
        onChange={async (e) => {
          uploadFile(e.target.files![0], {}, (p) => {}).then((v: any) => {
            // console.log(v.data.url);
            onChangeHeaderImageUrl?.(v.data.url);
          });
        }}
      />
    </div>
  );
};
export default WhatsappTemplateViewer;
