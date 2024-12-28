import { Card } from "@/components/ui/card";

export const PointsCard = () => {
  return (
    <Card className="w-full max-w-md p-6 bg-gradient-to-br from-primary-600 to-primary-800 text-white">
      <div className="space-y-4">
        <h3 className="text-lg font-medium opacity-90">Available Points</h3>
        <div className="flex items-baseline space-x-2">
          <span className="text-4xl font-bold">2,450</span>
          <span className="text-sm opacity-75">points</span>
        </div>
        <div className="text-sm opacity-75">
          Next reward available at 3,000 points
        </div>
        <div className="w-full bg-white/20 rounded-full h-2">
          <div className="bg-secondary w-4/5 h-2 rounded-full"></div>
        </div>
      </div>
    </Card>
  );
};