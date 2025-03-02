import { Button } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { verifyEmail } from "../services/api/authApi";

interface VerifyProps {}

const Verify: FC<VerifyProps> = ({}) => {
  const navigate = useNavigate();
  const [succeed, setSucceed] = useState(false);
  const { token } = useParams();
  const [error, setError] = useState("");
  const [mounted, setMounted] = useState();

  useEffect(() => {
    if (token) {
      const verifyToken = async () => {
        try {
          await verifyEmail(token);
          setSucceed(true);
        } catch (error) {
          setError(`${error}`);
          // alert(error);
        }
      };
      verifyToken();
    }
  }, []);
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 h-screen w-screen p-20">
     

      {succeed && (
        <div className="bg-white p-6 rounded shadow-md">
          <h1 className="text-2xl font-bold mb-4">Congratulations!</h1>
          <p className="text-lg mb-4">
            Your account has been successfully activated.
          </p>
          <Button
            onClick={() => {
              navigate("/login");
              // Navigate to login page
            }}
          >
            Go to Login
          </Button>
        </div>
      )}
    </div>
  );
};
export default Verify;
