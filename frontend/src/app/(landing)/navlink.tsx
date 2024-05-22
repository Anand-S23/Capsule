type NavLinkProps = {
    title: string;
    href: string;
};

export const NavLink = ({ title, href }: NavLinkProps) => {
    return (
        <p className="text-lg text-slate-600 hover:text-violet-600 hover:underline underline-offset-4 decoration-violet-600 hover:cursor-pointer">
            {title}
        </p>
    );
}

export const NavLinkButton = ({ title, href }: NavLinkProps) => {
    return (
        <p className="text-lg text-slate-600 bg-white hover:text-violet-600 hover:bg-violet-100 hover:cursor-pointer p-4 w-full flex justify-around">
            {title}
        </p>
    );
}

