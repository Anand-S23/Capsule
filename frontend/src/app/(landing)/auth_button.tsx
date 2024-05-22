'use client';

import { Button } from "@/components/ui/button";
import { AUTH_USER_ENDPOINT, LOGOUT_ENDPOINT } from "@/lib/consts";
import { LoaderIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export const AuthButton = () => {

    const router = useRouter();

    const [userID, setUserID] = useState<string>('');
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
                setUserID('');
            } else {
                const userID: string = await response.json();
                setUserID(userID);
            }

            setIsLoaded(true);
        }

        doAuth();
    }, []);

    const logout = async () => {
        await fetch(LOGOUT_ENDPOINT, {
            method: "POST",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
            credentials: 'include'
        });

        router.push('/');
    }

    if (!isLoaded) {
        return (
            <LoaderIcon w-8 h-8 className="animate-spin"/>
        );
    }

    return (
        <div>
            { userID === '' &&
                <div className="flex items-center justify-between gap-x-3">
                    <Button variant="outline" onClick={() => router.push("/login")}>
                        Login
                    </Button>

                    <Button onClick={() => router.push("/register")}>
                        Sign Up
                    </Button>
                </div>
            }

            { userID !== '' && 
                <></>
            }
        </div>
    );
}

