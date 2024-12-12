"use client";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import { useTheme } from "next-themes";
import Link from "next/link";

export const Hero = () => {
  return (
    <section className="flex flex-col justify-center items-center">
      <div className="grid place-content-center gap-8 py-20 md:py-32">
        <div className="text-center space-y-8">
          <Badge variant="outline" className="text-sm py-2">
            <span className="mr-2 text-primary">
              <Badge variant="destructive" className="text-white">BETA</Badge>
            </span>
            <span>FlowForge is now live!</span>
          </Badge>

          <div className="max-w-screen-md mx-auto text-center text-5xl font-bold">
            <h1>
              Transform Your
              <span className="text-transparent px-2 bg-gradient-to-r from-secondary to-destructive bg-clip-text">
                Codebase
              </span>
              Into Intuitive Flowcharts
            </h1>
          </div>

          <p className="max-w-screen-sm mx-auto text-xl text-muted-foreground">
            Simplify complex systems, visualize dependencies, and streamline
            your development process with FlowForge – the ultimate tool for
            developers.
          </p>

          <div className="space-y-4 md:space-y-0 md:space-x-4">
            <Button className="w-5/6 md:w-1/4 font-semibold text-white group/arrow">
              Try for Free
              <ArrowRight className="size-5 ml-2 group-hover/arrow:translate-x-1 transition-transform" />
            </Button>

            <Button
              asChild
              variant="secondary"
              className="w-5/6 md:w-1/4 font-semibold text-white"
            >
              <Link
                href="https://github.com/TheRSFoundations/flowforge"
                target="_blank"
              >
                View on GitHub
              </Link>
            </Button>
          </div>
        </div>
      </div>
    </section>
  );
};
