import { Button } from "flowbite-react";
import type { FC } from "react";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_TOKEN } from "../utils/constants";
import Logo from "../components/logo";
import { processLogin } from "../services/api/authApi";
import { LoadingContext } from "../contexts/LoadingContext";
import toast, { Toaster } from "react-hot-toast";

interface LoginProps {}

const Login: FC<LoginProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);

  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      setLoading(true);
      let resp: any = await processLogin({
        email,
        password,
      });

      await asyncStorage.setItem(LOCAL_STORAGE_TOKEN, resp.token);
      window.location.reload();
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="flex justify-center items-center h-screen bg-gray-50 dark:bg-gray-900 bg-no-repeat bg-cover"
      style={{
        backgroundSize: "cover",
        backgroundImage:
          'url("https://images.unsplash.com/photo-1738251396922-b6ef53f67b72")',
        backgroundPosition: "center",
      }}
    >
      <Toaster position="bottom-left" reverseOrder={false} />
      <section className="w-full px-10">
        <div
          className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0 "
          style={{ maxWidth: "500px" }}
        >
          <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
              <div className="flex justify-center mb-8">
                <Logo />
              </div>
              <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                Sign in to your account
              </h1>
              <form
                onSubmit={handleSubmit}
                className="space-y-4 md:space-y-6"
                action="#"
              >
                <div>
                  <label
                    htmlFor="email"
                    className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Your email
                  </label>
                  <input
                    type="email"
                    name="email"
                    id="email"
                    value={email}
                    onChange={(event) => setEmail(event.target.value)}
                    className="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="name@company.com"
                    required
                  />
                </div>
                <div>
                  <label
                    htmlFor="password"
                    className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Password
                  </label>
                  <input
                    type="password"
                    name="password"
                    id="password"
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                    placeholder="••••••••"
                    className="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    required
                  />
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-start">
                    <div className="flex items-center h-5">
                      <input
                        id="remember"
                        aria-describedby="remember"
                        type="checkbox"
                        className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800"
                      />
                    </div>
                    <div className="ml-3 text-sm">
                      <label
                        htmlFor="remember"
                        className="text-gray-500 dark:text-gray-300"
                      >
                        Remember me
                      </label>
                    </div>
                  </div>
                  <a
                    href="#"
                    className="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"
                  >
                    Forgot password?
                  </a>
                </div>
                <Button type="submit" className="w-full">
                  Sign in
                </Button>
                <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                  Don’t have an account yet?{" "}
                  <a
                    href="#"
                    className="font-medium text-primary-600 hover:underline dark:text-primary-500"
                    onClick={() => navigate("/register")}
                  >
                    Sign up
                  </a>
                </p>
              </form>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};
export default Login;
