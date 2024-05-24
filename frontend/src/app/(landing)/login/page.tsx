'use client';

import { useRouter } from "next/navigation";
import LoginForm from "./login-form";

const Login = () => {
    const router = useRouter();

    return (
        <div className="flex justify-center items-center h-full py-12">
            <div className="bg-white p-8 rounded shadow-lg w-96">
                <h1 className="px-5 text-2xl text-semibold pt-5"> Login </h1>

                <LoginForm />

                <p className="px-5 flex items-center justify-center">
                    New User? <span 
                        className="text-blue-500 hover:text-blue-400 hover:cursor-pointer hover:underline pl-1"
                        onClick={() => router.push('/register')}
                    >Sign Up</span>
                </p> 
            </div>
        </div>
    );
}

export default Login;

