import { Navbar } from "@/components/Navbar";
import { PointsCard } from "@/components/PointsCard";
import { RestaurantCard } from "@/components/RestaurantCard";
import { Button } from "@/components/ui/button";

const Index = () => {
  const featuredRestaurants = [
    {
      name: "Pasta Paradise",
      cuisine: "Italian",
      rating: 4.8,
      image: "https://images.unsplash.com/photo-1481931098730-318b6f776db0?auto=format&fit=crop&q=80",
      points: 3,
    },
    {
      name: "Sushi Master",
      cuisine: "Japanese",
      rating: 4.9,
      image: "https://images.unsplash.com/photo-1579871494447-9811cf80d66c?auto=format&fit=crop&q=80",
      points: 2,
    },
    {
      name: "Burger Bliss",
      cuisine: "American",
      rating: 4.7,
      image: "https://images.unsplash.com/photo-1551782450-17144efb9c50?auto=format&fit=crop&q=80",
      points: 2,
    },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      {/* Hero Section */}
      <section className="pt-32 pb-16 px-4">
        <div className="container mx-auto">
          <div className="flex flex-col md:flex-row items-center gap-12">
            <div className="flex-1 space-y-6">
              <h1 className="text-5xl font-bold text-gray-900">
                Earn rewards at your favorite restaurants
              </h1>
              <p className="text-xl text-gray-600">
                Join Payback and start earning points with every meal. Redeem for exclusive rewards and experiences.
              </p>
              <div className="flex gap-4">
                <Button size="lg" className="bg-primary-700 hover:bg-primary-800">
                  Get Started
                </Button>
                <Button size="lg" variant="outline">
                  Learn More
                </Button>
              </div>
            </div>
            <div className="flex-1 flex justify-center">
              <div className="animate-float">
                <PointsCard />
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Featured Restaurants */}
      <section className="py-16 px-4 bg-white">
        <div className="container mx-auto">
          <h2 className="text-3xl font-bold text-gray-900 mb-8">
            Featured Restaurants
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {featuredRestaurants.map((restaurant) => (
              <RestaurantCard key={restaurant.name} {...restaurant} />
            ))}
          </div>
        </div>
      </section>
    </div>
  );
};

export default Index;