import type { FC } from "react";
import { useParams } from "react-router-dom";

interface FormPublicPageProps {}

const FormPublicPage: FC<FormPublicPageProps> = ({}) => {
  const { formCode } = useParams();
  return <h1>{formCode}</h1>;
};
export default FormPublicPage;
