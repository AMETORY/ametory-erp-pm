import { useEffect, useState, type FC } from "react";
import { FormField, FormFieldType, FormSection } from "../models/form";
import {
  Button,
  Checkbox,
  Datepicker,
  FileInput,
  Label,
  Radio,
  Select,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import moment from "moment";

interface FormViewProps {
  sections: FormSection[];
  onSubmit: (val: FormSection[]) => void;
}

const FormView: FC<FormViewProps> = ({ sections, onSubmit }) => {
  const [sectionValues, setSectionValues] = useState([...sections]);

  useEffect(() => {
    if (sections.length) {
      setSectionValues([...sections]);
    }
  }, [sections]);
  const renderField = (
    sectionIndex: number,
    fieldIndex: number,
    field: FormField
  ) => {
    if (!sectionValues) return;
    if (sectionValues.length == 0) return;
    switch (field.type) {
      case FormFieldType.TextField:
        return (
          <div key={field.id}>
            <TextInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.PasswordField:
        return (
          <div key={field.id}>
            <TextInput
              type="password"
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.FileUpload:
        return (
          <div key={field.id}>
            <FileInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              onChange={(e) => {
                const file = e.target.files;
                if (file) {
                  const reader = new FileReader();
                  reader.onload = (event) => {
                    sectionValues[sectionIndex].fields[fieldIndex].value = event
                      .target?.result as string;
                    setSectionValues(sectionValues);
                  };
                  reader.readAsDataURL(file[0]);
                }

                // sectionValues[sectionIndex].fields[fieldIndex].value = e.target.value
                // setSectionValues(sectionValues)
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.EmailField:
        return (
          <div key={field.id}>
            <TextInput
              type={"email"}
              sizing={"sm"}
              name={field.label}
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.NumberField:
        return (
          <div key={field.id}>
            <TextInput
              type={"number"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.Currency:
        return (
          <div key={field.id}>
            <TextInput
              type={"number"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.TextArea:
        return (
          <div key={field.id}>
            <Textarea
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  e.target.value;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DatePicker:
        return (
          <div key={field.id}>
            <Datepicker
              placeholder={field.placeholder}
              required={field.required}
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(e) => {
                sectionValues[sectionIndex].fields[fieldIndex].value = e;
                setSectionValues(sectionValues);
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DateRangePicker:
        if (!sectionValues[sectionIndex].fields[fieldIndex].value) {
          sectionValues[sectionIndex].fields[fieldIndex].value = [
            new Date(),
            new Date(),
          ];
        }
        return (
          <div className="">
            <div className="grid grid-cols-2 gap-2">
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
                value={moment(
                  sectionValues[sectionIndex].fields[fieldIndex].value[0]
                ).toDate()}
                onChange={(e) => {
                  let val = [
                    ...sectionValues[sectionIndex].fields[fieldIndex].value,
                  ];
                  setSectionValues([
                    ...sectionValues.map((section, i) => {
                      if (i == sectionIndex) {
                        return {
                          ...section,
                          fields: [
                            ...section.fields.map((field, j) => {
                              if (j == fieldIndex) {
                                return {
                                  ...field,
                                  value: [e, val[1]],
                                };
                              }
                              return field;
                            }),
                          ],
                        };
                      }
                      return section;
                    }),
                  ]);
                }}
              />
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
                value={moment(
                  sectionValues[sectionIndex].fields[fieldIndex].value[1]
                ).toDate()}
                onChange={(e) => {
                  let val = [
                    ...sectionValues[sectionIndex].fields[fieldIndex].value,
                  ];

                  setSectionValues([
                    ...sectionValues.map((section, i) => {
                      if (i == sectionIndex) {
                        return {
                          ...section,
                          fields: [
                            ...section.fields.map((field, j) => {
                              if (j == fieldIndex) {
                                return {
                                  ...field,
                                  value: [val[0], e],
                                };
                              }
                              return field;
                            }),
                          ],
                        };
                      }
                      return section;
                    }),
                  ]);
                }}
              />
            </div>
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.RadioButton:
        return (
          <div key={field.id}>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}
              {field.options.map((option, i) => (
                <div className="flex items-center gap-2" key={i}>
                  <Radio
                    id={`${field.label}-${i}`}
                    value={option.value}
                    checked={
                      sectionValues[sectionIndex].fields[fieldIndex].value ==
                      option.value
                    }
                    onChange={(val) => {
                      // sectionValues[sectionIndex].fields[fieldIndex].value =
                      //   val.target.value;
                      setSectionValues([
                        ...sectionValues.map((section, i) => {
                          if (i == sectionIndex) {
                            return {
                              ...section,
                              fields: [
                                ...section.fields.map((field, j) => {
                                  if (j == fieldIndex) {
                                    return {
                                      ...field,
                                      value: val.target.value,
                                    };
                                  }
                                  return field;
                                }),
                              ],
                            };
                          }
                          return section;
                        }),
                      ]);
                    }}
                  />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.Checkbox:
        if (!sectionValues[sectionIndex].fields[fieldIndex].value) {
          sectionValues[sectionIndex].fields[fieldIndex].value = [];
        }
        return (
          <div key={field.id}>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}

              {field.options.map((option, i) => (
                <div className="flex items-center gap-2" key={i}>
                  <Checkbox
                    id={`${field.label}-${i}`}
                    value={option.value}
                    checked={sectionValues[sectionIndex].fields[
                      fieldIndex
                    ].value.includes(option.value)}
                    onChange={(val) => {
                      if (
                        !sectionValues[sectionIndex].fields[
                          fieldIndex
                        ].value.includes(option.value)
                      ) {
                        sectionValues[sectionIndex].fields[
                          fieldIndex
                        ].value.push(option.value);
                      } else {
                        sectionValues[sectionIndex].fields[fieldIndex].value = [
                          ...sectionValues[sectionIndex].fields[
                            fieldIndex
                          ].value.filter((val: any) => val != option.value),
                        ];
                      }

                      setSectionValues([...sectionValues]);
                    }}
                  />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.ToggleSwitch:
        if (!sectionValues[sectionIndex].fields[fieldIndex].value) {
          sectionValues[sectionIndex].fields[fieldIndex].value = false;
        }
        return (
          <div key={field.id}>
            <ToggleSwitch
              sizing="sm"
              checked={sectionValues[sectionIndex].fields[fieldIndex].value}
              label={field.help_text}
              onChange={(val) => {
                sectionValues[sectionIndex].fields[fieldIndex].value = val;
                setSectionValues([
                  ...sectionValues.map((section, i) => {
                    if (i == sectionIndex) {
                      return {
                        ...section,
                        fields: [
                          ...section.fields.map((field, j) => {
                            if (j == fieldIndex) {
                              return {
                                ...field,
                                value: val,
                              };
                            }
                            return field;
                          }),
                        ],
                      };
                    }
                    return section;
                  }),
                ]);
              }}
            />
          </div>
        );
      case FormFieldType.Dropdown:
        return (
          <div key={field.id}>
            <Select
              value={sectionValues[sectionIndex].fields[fieldIndex].value}
              onChange={(val) => {
                sectionValues[sectionIndex].fields[fieldIndex].value =
                  val.target.value;
                setSectionValues(sectionValues);
              }}
            >
              {field.options.map((option, i) => (
                <option
                  selected={
                    sectionValues[sectionIndex].fields[fieldIndex].value ==
                    option.value
                  }
                  key={i}
                  value={option.value}
                >
                  {option.label}
                </option>
              ))}
            </Select>
          </div>
        );
    }

    return <div></div>;
  };
  return (
    <div className="flex flex-col justify-center w-1/2 space-y-4 ">
      {sections.map((section, index) => (
        <div
          className="bg-white p-4 rounded-lg border border-t-4 border-t-blue-400"
          key={section.id}
        >
          <h1 className="text-2xl font-bold">{section.section_title}</h1>
          <div className="text-md text-gray-600">{section?.description}</div>
          <div className="flex flex-col space-y-4 mt-4">
            {section.fields.map((field, fieldIndex) => (
              <div key={field.id}>
                <Label className="font-bold text-md">{field.label}</Label>
                {renderField(index, fieldIndex, field)}
              </div>
            ))}
          </div>
        </div>
      ))}
      <Button
        type="submit"
        onClick={() => {
          onSubmit(sectionValues);
        }}
      >
        Submit Form
      </Button>
    </div>
  );
};
export default FormView;
