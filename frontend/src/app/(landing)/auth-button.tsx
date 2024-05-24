'use client';

import { Button } from "@/components/ui/button";
import { AUTH_USER_ENDPOINT } from "@/lib/consts";
import { LoaderIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { ProfileDropdown } from "./profile-dropdown";

interface AuthUser {
    ID: string;
    Name: string;
}

export const AuthButton = () => {
    const router = useRouter();

    const [user, setUser] = useState<AuthUser>({ ID: '', Name: ''});
    const [isLoaded, setIsLoaded] = useState<boolean>(false);

    useEffect(() => {
        const doAuth = async () => {
            const response = await fetch(AUTH_USER_ENDPOINT, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
                credentials: 'include'
            });

            if (!response.ok) {
                setUser({ ID: '', Name: '' });
            } else {
                const authUser: AuthUser = await response.json() as AuthUser;
                console.log(authUser);
                setUser(authUser);
            }

            setIsLoaded(true);
        }

        doAuth();
    }, []);

    if (!isLoaded) {
        return (
            <LoaderIcon className="animate-spin"/>
        );
    }

    return (
        <div>
            { user.ID === '' &&
                <div className="flex items-center justify-between gap-x-3">
                    <Button variant="outline" onClick={() => router.push("/login")}>
                        Login
                    </Button>

                    <Button onClick={() => router.push("/register")}>
                        Sign Up
                    </Button>
                </div>
            }

            { user.ID !== '' && 
                <ProfileDropdown username={user.Name} />
            }
        </div>
    );
}

