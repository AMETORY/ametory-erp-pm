import { useEffect, useState, type FC } from "react";
import { FormTemplateModel } from "../models/form";
import { Button, Select, TextInput } from "flowbite-react";

interface FormStyleProps {
  formTemplate: FormTemplateModel;
  style: any | null;
  onSave?: (d: FormTemplateModel) => void;
}

const FormStyle: FC<FormStyleProps> = ({ formTemplate, style, onSave }) => {
  const [formData, setFormData] = useState<any>({
    style: {
      backgroundColor: "#f3f4f6",
      textColor: "#333333",
      buttonColor: "#00b7e9",
      fontSize: 14,
      border: 2,
    },
  });

  useEffect(() => {
    if (style) {
      setFormData({
        style,
      });
    }
  }, [style]);
  const handleStyleChange = (key: string, value: string | number) => {
    setFormData((prevData: any) => ({
      ...prevData,
      style: {
        ...prevData.style,
        [key]: value,
      },
    }));
  };

  const handleStyleSave = () => {
    onSave?.({
      ...formTemplate,
      style: formData.style,
    });
  };
  return (
    <>
      <div style={{ display: "flex", flexDirection: "column", gap: "15px" }}>
        <div className="grid grid-cols-2 gap-4">
          <div className="flex flex-col space-y-4">
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Cover Height (px):
              </label>
              <TextInput
                type="number"
                value={parseInt(formData?.style.coverHeight)}
                onChange={(e) =>
                  handleStyleChange("coverHeight", `${e.target.value}px`)
                }
                min="12"
                max="24"
                style={{ padding: "5px" }}
              />
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Font Label Size (px):
              </label>
              <TextInput
                type="number"
                value={parseInt(formData?.style.fontLabelSize)}
                onChange={(e) =>
                  handleStyleChange("fontLabelSize", `${e.target.value}px`)
                }
                min="12"
                max="24"
                style={{ padding: "5px" }}
              />
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Font Size (px):
              </label>
              <TextInput
                type="number"
                value={parseInt(formData?.style.fontSize)}
                onChange={(e) =>
                  handleStyleChange("fontSize", `${e.target.value}px`)
                }
                min="12"
                max="24"
                style={{ padding: "5px" }}
              />
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Font Family:
              </label>
              <Select
                className="select"
                value={formData?.style.fontFamily}
                onChange={(e) =>
                  handleStyleChange("fontFamily", e.target.value)
                }
                style={{ padding: "5px", width: "100%" }}
              >
                <option value="inherit">Default</option>
                <option value="Arial, sans-serif">Arial</option>
                <option value="Georgia, serif">Georgia</option>
                <option value="Times New Roman, serif">Times New Roman</option>
                <option value="Courier New, monospace">Courier New</option>
                <option value="Segoe UI, Tahoma, Geneva, Verdana, sans-serif">
                  Segoe UI
                </option>
                <option value="Helvetica, Arial, sans-serif">Helvetica</option>
              </Select>
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Border:
              </label>
              <TextInput
                type="number"
                min="0"
                max="5"
                value={parseInt(formData?.style.border ?? 0)}
                onChange={(e) => handleStyleChange("border", e.target.value)}
                style={{ padding: "5px", width: "100%" }}
                placeholder="e.g., 2px solid #333"
              />
            </div>

            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Box Shadow:
              </label>
              <TextInput
                type="text"
                value={formData?.style.boxShadow}
                onChange={(e) => handleStyleChange("boxShadow", e.target.value)}
                style={{ padding: "5px", width: "100%" }}
                placeholder="e.g., 5px 5px 15px rgba(0,0,0,0.3)"
              />
            </div>
          </div>
          <div className="flex flex-col space-y-4">
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Background Color:
              </label>
              <input
                type="color"
                className="input"
                value={formData?.style.backgroundColor}
                onChange={(e) =>
                  handleStyleChange("backgroundColor", e.target.value)
                }
              />
            </div>

            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Text Color:
              </label>
              <input
                type="color"
                className="input"
                value={formData?.style.textColor}
                onChange={(e) => handleStyleChange("textColor", e.target.value)}
              />
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Border Color:
              </label>
              <input
                type="color"
                className="input"
                value={formData?.style.borderColor}
                onChange={(e) =>
                  handleStyleChange("borderColor", e.target.value)
                }
              />
            </div>
            <div>
              <label style={{ display: "block", marginBottom: "5px" }}>
                Button Color:
              </label>
              <input
                type="color"
                className="input"
                value={formData?.style.buttonColor}
                onChange={(e) =>
                  handleStyleChange("buttonColor", e.target.value)
                }
              />
            </div>
          </div>
        </div>
      </div>
      <div style={{ marginTop: "20px" }}>
        <Button onClick={handleStyleSave}>Save Changes</Button>
      </div>
    </>
  );
};
export default FormStyle;
