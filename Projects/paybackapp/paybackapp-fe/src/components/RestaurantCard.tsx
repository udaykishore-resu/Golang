import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

interface RestaurantCardProps {
  name: string;
  cuisine: string;
  rating: number;
  image: string;
  points: number;
}

export const RestaurantCard = ({
  name,
  cuisine,
  rating,
  image,
  points,
}: RestaurantCardProps) => {
  return (
    <Card className="overflow-hidden group cursor-pointer transition-all hover:shadow-lg">
      <div className="relative h-48 overflow-hidden">
        <img
          src={image}
          alt={name}
          className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        />
        <Badge className="absolute top-3 right-3 bg-secondary">
          {points}x Points
        </Badge>
      </div>
      <div className="p-4">
        <h3 className="font-semibold text-lg">{name}</h3>
        <p className="text-gray-600 text-sm">{cuisine}</p>
        <div className="flex items-center mt-2">
          <span className="text-yellow-400">★</span>
          <span className="ml-1 text-sm text-gray-600">{rating}</span>
        </div>
      </div>
    </Card>
  );
};