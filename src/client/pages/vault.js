import React from "react";
import Head from "next/head";
import Nav from "@/components/Nav";
import SiteNav from "@/components/SiteNav/SiteNav";
import Files from "@/components/Files/Files";

const Vault = () => {
    return (
        <main className="md:grid md:grid-cols-12">
            <Head>
                <title>Vault</title>
            </Head>
            <SiteNav />
            <div className="md:col-start-4 md:col-span-full lg:col-start-3">
                <Nav />
                <Files />
            </div>
        </main>
    );
};

export default Vault;
