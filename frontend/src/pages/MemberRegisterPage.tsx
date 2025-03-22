import { useContext, useEffect, useState, type FC } from "react";
import { useParams } from "react-router-dom";
import { LoadingContext } from "../contexts/LoadingContext";
import { createContact } from "../services/api/contactApi";
import {
  getNumber,
  whatsappSessionAuthRegister,
} from "../services/api/commonApi";
import toast, { Toaster } from "react-hot-toast";

interface MemberRegisterPageProps {}

const MemberRegisterPage: FC<MemberRegisterPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    address: "",
  });

  const { code } = useParams();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };
  const handleChangeAddress = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    let dataRegis = {
      name: formData.name,
      email: formData.email,
      phone: formData.phone,
      address: formData.address,
      code: code,
    };
    try {
      setLoading(true);
      await whatsappSessionAuthRegister(dataRegis);
      toast.success("Register Success");
      setTimeout(() => {
        window.close();
      }, 3000);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    setLoading(true);
    getNumber(code!)
      .then((resp: any) => {
        setFormData({
          ...formData,
          phone: resp.data,
        });
      })
      .catch((error: any) => {
        toast.error(`${error}`);
      })
      .finally(() => {
        setLoading(false);
      });
  }, [code]);

  return (
    <div className=" w-full h-screen mx-auto flex items-center justify-center bg-gray-400">
      <Toaster position="bottom-left" reverseOrder={false} />
      <div className="bg-white rounded-lg shadow-lg p-8 max-w-[600px] w-full mt-20">
        <h1 className="text-2xl font-bold mb-4">Registration Form</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block font-semibold mb-1" htmlFor="name">
              Name
            </label>
            <input
              type="text"
              id="name"
              name="name"
              value={formData.name}
              onChange={handleChange}
              className="border border-gray-200 rounded p-2 w-full"
              placeholder="Enter your name"
              required
            />
          </div>
          <div>
            <label className="block font-semibold mb-1" htmlFor="email">
              Email
            </label>
            <input
              type="email"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              className="border border-gray-200 rounded p-2 w-full"
              placeholder="Enter your email"
              required
            />
          </div>
          <div>
            <label className="block font-semibold mb-1" htmlFor="phone">
              Phone
            </label>
            <input
              readOnly
              type="tel"
              id="phone"
              name="phone"
              value={formData.phone}
              onChange={handleChange}
              className="border border-gray-200 rounded p-2 w-full"
              placeholder="Enter your phone number"
              required
            />
          </div>
          <div>
            <label className="block font-semibold mb-1" htmlFor="address">
              Address
            </label>
            <textarea
              id="address"
              name="address"
              value={formData.address}
              onChange={handleChangeAddress}
              className="border border-gray-200 rounded p-2 w-full"
              placeholder="Enter your address"
              required
              rows={5}
            />
          </div>
          <button type="submit" className="bg-blue-500 text-white rounded p-2">
            Register
          </button>
        </form>
      </div>
    </div>
  );
};
export default MemberRegisterPage;
