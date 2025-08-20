import { useContext, useEffect, useState, type FC } from "react";
import { useParams } from "react-router-dom";
import { getFormPublic, postFormPublic } from "../services/api/formApi";
import { FormModel } from "../models/form";
import FormView from "../components/FormView";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";

interface FormPublicPageProps {}

const FormPublicPage: FC<FormPublicPageProps> = ({}) => {
  const { formCode } = useParams();
  const [form, setForm] = useState<FormModel>();
  const { loading, setLoading } = useContext(LoadingContext);

  useEffect(() => {
    if (formCode) {
      getFormPublic(formCode).then((e: any) => setForm(e.data));
    }
  }, [formCode]);

  const handleSubmitForm = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
  };

  return (
    <form
      onSubmit={handleSubmitForm}
      className="bg-gray-50 flex flex-col  items-center p-16 overflow-y-auto h-[calc(100vh)] "
      style={{
        backgroundColor: form?.form_template?.style?.backgroundColor,
      }}
    >
      <div className="flex flex-col justify-center w-1/2 space-y-4 mb-4">
        <div
          className="bg-white rounded-lg border border-t-4 border-t-blue-400"
          style={{
            borderColor: form?.form_template?.style?.borderColor,
            borderTopWidth: form?.form_template?.style?.border
              ? `${form?.form_template?.style?.border}px`
              : 0,
          }}
        >
          {form?.cover && (
            <img
              src={form?.cover?.url}
              className=" aspect-video w-full object-cover"
              style={{
                height: form?.form_template?.style?.coverHeight,
              }}
            />
          )}
          <div className="p-4  ">
            <h1
              className="text-2xl font-semibold"
              style={{
                fontFamily: form?.form_template?.style?.fontFamily,
                color: form?.form_template?.style?.textColor,
              }}
            >
              {form?.title}
            </h1>
            <p
              style={{
                fontFamily: form?.form_template?.style?.fontFamily,
                color: form?.form_template?.style?.textColor,
              }}
            >
              {form?.description}
            </p>
          </div>
        </div>
      </div>
      {form && (
        <FormView
          style={form?.form_template?.style}
          sections={form?.form_template?.sections ?? []}
          onSubmit={async (val) => {
            try {
              setLoading(true);
              await postFormPublic(formCode!, val);
              if (window.parent !== window) {
                window.parent.postMessage(
                  JSON.stringify({
                    type: "FORM_SUBMITTED",
                    data: {
                      formCode: formCode,
                    },
                  }),
                  "*"
                );
              }
              alert("Form submitted successfully");
            } catch (error) {
              toast.error(`${error}`);
            } finally {
              setLoading(false);
            }
          }}
        />
      )}
    </form>
  );
};
export default FormPublicPage;
