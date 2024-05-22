import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";

const font = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Capsule",
    description: "Cultivate your professional network into meaningful relationships",
};

export default function RootLayout({ children }: Readonly<{children: React.ReactNode}>) {
    return (
        <html lang="en">
            <body className={cn(font.className, "overflow-y-hidden")}>
                {children}
            </body>
        </html>
    );
}
