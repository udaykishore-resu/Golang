import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export const Navbar = () => {
  const navigate = useNavigate();

  return (
    <nav className="fixed top-0 w-full bg-white/80 backdrop-blur-md z-50 border-b">
      <div className="container mx-auto px-4 h-16 flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <h1 className="text-2xl font-bold text-primary-700">Payback</h1>
        </div>
        <div className="flex items-center space-x-4">
          <Button
            variant="ghost"
            onClick={() => navigate("/restaurants")}
            className="text-gray-600 hover:text-primary-700"
          >
            Restaurants
          </Button>
          <Button
            variant="ghost"
            onClick={() => navigate("/rewards")}
            className="text-gray-600 hover:text-primary-700"
          >
            Rewards
          </Button>
          <Button
            variant="ghost"
            onClick={() => navigate("/profile")}
            className="text-gray-600 hover:text-primary-700"
          >
            Profile
          </Button>
        </div>
      </div>
    </nav>
  );
};