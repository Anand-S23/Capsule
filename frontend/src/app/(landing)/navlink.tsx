import { useRouter } from "next/navigation";

type NavLinkProps = {
    title: string;
    href: string;
};

export const NavLink = ({ title, href }: NavLinkProps) => {
    const router = useRouter();

    return (
        <p className="text-md text-slate-700 hover:text-violet-600 hover:underline underline-offset-4 decoration-violet-600 hover:cursor-pointer" onClick={() => router.push(href)}>
            {title}
        </p>
    );
}

export const NavLinkButton = ({ title, href }: NavLinkProps) => {
    const router = useRouter();

    return (
        <p className="text-md text-slate-700 bg-white hover:text-violet-600 hover:bg-violet-100 hover:cursor-pointer p-4 w-full flex justify-around">
            {title}
        </p>
    );
}

