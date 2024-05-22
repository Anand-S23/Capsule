'use client';

import { MenuIcon, XIcon } from "lucide-react";
import { AuthButton } from "./auth_button";
import { NavLink, NavLinkButton } from "./navlink";
import { useState } from "react";

export type LinkType = {
    title: string;
    href: string;
}

const links: Array<LinkType> = [
    { title: "Test 0", href: "/#hero" },
    { title: "Test 1", href: "/#about" },
    { title: "Test 2", href: "/#testimonials" },
    { title: "Test 3", href: "/#faq" },
    { title: "Test 4", href: "/#pricing" },
]

export const Header = () => {
    const [open, setOpen] = useState<boolean>(false);

    return (
        <header className="h-20 w-full border-b-2 border-slate-200">
            <nav className="lg:max-w-screen-2xl mx-auto flex items-center justify-between h-full px-4">
                <div className="pt-8 pl-4 pb-7 flex items-center gap-x-3">
                    <button className="bg-violet-600 h-10 w-10"></button>
                    <h1 className="text-2xl font-extrabold hidden md:block">
                        Capsule
                    </h1>
                </div>

                <div className="pt-8 pb-7 lg:flex hidden items-center justify-between gap-x-8">
                    { links.map((link: LinkType, idx: number) =>
                            <NavLink 
                                key={idx}
                                title={link.title} 
                                href={link.href}
                            />
                    )}
                </div>

                <div className="flex items-center justify-between">
                    <AuthButton />

                    <button className="px-4 lg:hidden" onClick={() => setOpen(!open)}>
                        {open && <XIcon w-8 h-8 />}
                        {!open && <MenuIcon w-8 h-8 />}
                    </button>
                </div>
            </nav>

            {open &&
                <div className="relative px-1">
                    <div className="mx-auto z-10 border-solid border border-slate-200 rounded-md shadow-md w-full flex flex-col items-center lg:hidden">
                        { links.map((link: LinkType, idx: number) =>
                                <NavLinkButton
                                    key={idx}
                                    title={link.title} 
                                    href={link.href}
                                />
                        )}
                    </div>
                </div>
            }
        </header>
    );
}
