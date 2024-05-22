import { Footer } from "./footer";
import { Header } from "./header";

type LandingPageProps = {
    children: React.ReactNode;
};

const LandingPageLayout = ({ children }: LandingPageProps) => {
    return (
        <div className="min-h-screen flex flex-col">
            <Header />

            <main className="max-w-screen-2xl my-auto">
                {children}
            </main>
            
            <Footer />
        </div>
    );
}

export default LandingPageLayout;
