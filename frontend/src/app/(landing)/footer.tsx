'use client';

import { useRouter } from "next/navigation";
import { NavLink } from "./navlink";

export const Footer = () => {
    const router = useRouter();

    return (
        <footer className="h-20 w-full border-t border-slate-200 p-2">
            <div className="lg:max-w-screen-2xl mx-auto flex items-center justify-between p-4">
                <div className="flex justify-between text-lg font-extrabold gap-x-2 hover:cursor-pointer" onClick={() => router.push('/')}>
                    <button className="h-8 w-8 bg-violet-600"></button>
                    Capsule
                </div>

                <div className="hidden sm:flex items-center justify-between gap-x-3">
                    <NavLink title="Terms of Service" href="/tos" />
                    <NavLink title="Privacy Policy" href="/privacy-policy" />
                </div>

                <div className="text-slate-600">
                    © 2024 Anand Singh
                </div>
            </div>
        </footer>
    );
}
