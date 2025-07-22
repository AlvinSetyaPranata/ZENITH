import { MetaProvider, Title } from "@solidjs/meta";
import { useNavigate } from "@solidjs/router";
import {
  createEffect,
  createResource,
  createSignal,
  ErrorBoundary,
  onMount,
  Show,
} from "solid-js";
import Card from "~/components/atoms/cards/card";
import CardWithStatus from "~/components/atoms/cards/cardWithStatus";
import DashboardLayout from "~/components/layouts/DashboardLayout";
import RootLayout from "~/components/layouts/DashboardLayout";
import { DefaultStudent } from "~/constants/student-constant";
import { GetStudentData } from "~/services/api/students/get-student-profile";
import { StudentType } from "~/types/master-types/profile-type";

export default function Home() {
  const navigate = useNavigate();
  const [data, { refetch }] = createResource(GetStudentData);

  createEffect(() => {
    if (data.error) {
      navigate("/");
    }
  }, [data.error]);

  createEffect(() => {
    if (data() == false) {
      navigate("/")
    }
  }, [data])

  onMount(() => {
    refetch();
  });

  return (
    <DashboardLayout>
      <MetaProvider>
        <Title>Zenith - Dashboard</Title>
      </MetaProvider>

      {/* content */}
      <main class="text-left mx-auto text-gray-700 p-8 bg-background-800 min-h-screen">
        <Show when={data()?.name} fallback={<div class="h-[20px] w-[20%] bg-gray-600 animate-pulse "></div>}>
          <h2 class="text-2xl text-white font-medium">
            Selamat pagi, {data().name}!
          </h2>
        </Show>

        <p class="mt-2 text-gray-400">
          Berikut adalah dashboard pribadi mu, {`(p>_<p)`}
        </p>

        <section class="w-full grid grid-cols-3 mt-14">
          <Card title="Kelas" icon="charm:graduate-cap" desc="1" />
          <Card title="Tugas" icon="material-symbols:task-outline" desc="1" />
          <Card
            title="Ipk"
            icon="material-symbols-light:stars-outline"
            desc="1"
          />
        </section>

        <section class="mt-14">
          <h2 class="text-2xl text-white font-medium mt-24">
            Aktivitas Perkuliahan
          </h2>
          <div class="w-full grid grid-cols-3 mt-12">
            <CardWithStatus
              title="Pemograman Dasar"
              time="0000-01-01T12:00:00Z"
              status="Berlangsung"
            />
            <Card
              title="Mata kuliah hari ini"
              icon="charm:graduate-cap"
              desc="4"
            />
            <Card
              title="Mata kuliah dihadiri"
              icon="mingcute:time-line"
              desc="4"
            />
          </div>
        </section>
      </main>
    </DashboardLayout>
  );
}
