"use client"

import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts"
import {
  ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
} from "@/components/ui/chart"
import { Locale } from "@/i18n.config";

interface ReviewsChartProps {
  locale: Locale;
  reviewsData: { name: string; count: number }[];
}

const chartConfig = {
  count: {
    label: "Reviews",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig

const ChartTooltipContent = ({ active, payload }: any) => {
  if (active && payload && payload.length) {
    const data = payload[0].payload;
    return (
      <div className="bg-background border border-secondary ring-2 ring-primary/50 rounded-md p-2 shadow-lg">
        <p className="font-medium text-foreground">{data.name}</p>
        <p className="text-sm text-muted-foreground">
          Reviews: <span className="font-medium text-primary">{data.count}</span>
        </p>
      </div>
    );
  }
  return null;
}

export default function ReviewsChart({ reviewsData }: ReviewsChartProps) {
  const chartData = reviewsData.map((item) => ({
    name: item.name,
    count: item.count,
  }));

  return (
    <div className="flex flex-col items-center justify-center gap-4 w-full">
      <p className="font-bold text-primary text-xl">Reviews by User</p>
      <ChartContainer config={chartConfig} className="p-0 w-full">
        <BarChart data={chartData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
          <CartesianGrid strokeDasharray="3 3" stroke="#ddd" />

          {/* X Axis */}
          <XAxis
            dataKey="name"
            tickLine={false}
            tickMargin={15}
            axisLine={false}
            tickFormatter={(value) => value.slice(0, 3)}
            style={{ fontSize: "12px", fill: "#6b7280" }}
          />

          {/* Y Axis */}
          <YAxis
            tickLine={false}
            axisLine={false}
            tickFormatter={(value) => value}
            style={{ fontSize: "12px", fill: "#6b7280" }}
          />

          <ChartTooltip content={<ChartTooltipContent />} />

          <ChartLegend content={<ChartLegendContent />} />

          {/* Bar */}
          <Bar
            dataKey="count"
            fill="var(--color-desktop)"
            radius={3}
            barSize={20}
            opacity={0.8}
            isAnimationActive={true}
            animationDuration={1300}
            animationEasing="ease-out"
          />
        </BarChart>
      </ChartContainer>
    </div>
  );
}
