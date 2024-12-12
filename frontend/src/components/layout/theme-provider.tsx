"use client";

import * as React from "react";
import {
  ThemeProvider as NextThemesProvider,
  ThemeProviderProps,
} from "next-themes";

interface NewThemeProviderProps extends ThemeProviderProps {
  children: React.ReactNode;
}

export function ThemeProvider({ children, ...props }: NewThemeProviderProps) {
  return <NextThemesProvider {...props}>{children}</NextThemesProvider>;
}
