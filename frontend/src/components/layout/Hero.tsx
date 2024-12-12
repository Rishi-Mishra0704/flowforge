"use client";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import { useTheme } from "next-themes";
import Image from "next/image";
import Link from "next/link";

export const Hero = () => {
  const { theme } = useTheme();
  return (
    <section className="flex flex-col justify-center items-center">
      <div className="grid place-content-center gap-8 py-20 md:py-32">
        <div className="text-center space-y-8">
          <Badge variant="outline" className="text-sm py-2">
            <span className="mr-2 text-primary">
              <Badge variant="destructive" className="text-white">
                BETA
              </Badge>
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
            <Button variant="accent" className="w-5/6 md:w-1/4 font-semibold text-white group/arrow">
              <Link href="/flowchart" className="flex justify-center items-center cursor-pointer">
                Try for Free
                <ArrowRight className="size-5 ml-2 group-hover/arrow:translate-x-1 transition-transform" />
              </Link>
            </Button>

            <Button
              asChild
              variant="default"
              className="w-5/6 md:w-1/4 font-semibold text-white cursor-pointer"
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

      <div className="relative group mt-14">
        <div className="absolute top-2 lg:-top-8 left-1/2 transform -translate-x-1/2 w-[90%] mx-auto h-24 lg:h-80 bg-secondary/50 rounded-full blur-3xl"></div>
        <Image
          width={1200}
          height={1200}
          className="w-full md:w-[1200px] mx-auto rounded-lg relative rouded-lg leading-none flex items-center border border-t-2 border-secondary  border-t-secondary/30"
          src={theme === "dark" ? "/images/hero-dark" : "/images/hero-light"}
          alt="dashboard"
        />

        <div className="absolute bottom-0 left-0 w-full h-20 md:h-28 bg-gradient-to-b from-background/0 via-background/50 to-background rounded-lg"></div>
      </div>
    </section>
  );
};
