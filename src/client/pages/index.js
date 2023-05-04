import Link from "next/link";

export default function Home() {
    return (
        <>
            <div>Hello, I&apos;m building Vaulter</div>
            <Link href="/vault">Vault</Link>
        </>
    );
}
