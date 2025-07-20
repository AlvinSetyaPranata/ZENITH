import { MetaProvider, Title } from "@solidjs/meta";
import RootLayout from "~/components/layouts/RootLayout";

export default function PersonalDataPage() {
  return (
    <>
       <MetaProvider>
              <Title>Zenith - Data Diri Mahasiswa</Title>
            </MetaProvider>

      <main class="text-left mx-auto text-gray-700 p-8 bg-background-800 min-h-screen">

          <h1>This is biodata</h1>
      </main>
    </>
  )
}
