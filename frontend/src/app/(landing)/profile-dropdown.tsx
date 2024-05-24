
import {
    ChevronDown,
  CreditCard,
  Github,
  LifeBuoy,
  LogOut,
  Settings,
  User,
} from "lucide-react"

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { LOGIN_ENDPOINT } from "@/lib/consts";
import { useRouter } from "next/navigation";

type ProfileDropdownProps = {
    username: string;
}

export const ProfileDropdown = ({ username }: ProfileDropdownProps) => {
    const router = useRouter();

    const logout = async () => {
        await fetch(LOGIN_ENDPOINT, {
            method: "POST",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        router.push('/');
    }

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <div className="flex items-center justify-between gap-x-1 hover:cursor-pointer hover:text-slate-700">
                    <p className="text-lg">{username}</p>
                    <ChevronDown />
                </div>
            </DropdownMenuTrigger>

            <DropdownMenuContent className="w-56 cursor-default">
                <DropdownMenuLabel>My Account</DropdownMenuLabel>
                <DropdownMenuSeparator />

                <DropdownMenuGroup>
                    <DropdownMenuItem>
                        <User className="mr-2 h-4 w-4" />
                        <span>Profile</span>
                    </DropdownMenuItem>

                    <DropdownMenuItem>
                        <CreditCard className="mr-2 h-4 w-4" />
                        <span>Billing</span>
                    </DropdownMenuItem>

                    <DropdownMenuItem>
                        <Settings className="mr-2 h-4 w-4" />
                        <span>Settings</span>
                    </DropdownMenuItem>
                </DropdownMenuGroup>

                <DropdownMenuSeparator />

                <DropdownMenuItem>
                    <Github className="mr-2 h-4 w-4" />
                    <span>GitHub</span>
                </DropdownMenuItem>

                <DropdownMenuItem>
                    <LifeBuoy className="mr-2 h-4 w-4" />
                    <span>Support</span>
                </DropdownMenuItem>

                <DropdownMenuSeparator />

                <DropdownMenuItem onClick={logout}>
                    <LogOut className="mr-2 h-4 w-4" />
                    <span>Log out</span>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
