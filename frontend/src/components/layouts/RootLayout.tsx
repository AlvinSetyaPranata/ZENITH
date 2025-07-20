import {
  createEffect,
  createResource,
  createSignal,
  onMount,
  ParentProps,
  Show,
  useContext,
} from "solid-js";
import Sidebar from "../ui/sidebar";
import SidebarNav from "../molecules/sidebar-nav";
import { A, useLocation, useNavigate } from "@solidjs/router";
import { Icon } from "@iconify-icon/solid";
import { AuthContext } from "~/contexts/auth-context";
import { UserType } from "~/types/master-types/role-type";
import { AuthStoreType } from "~/types/auth-types/store";
import { GetStudentData } from "~/services/api/students/get-student-profile";
import { SoonerContext } from "~/contexts/sooner-context";

export default function RootLayout({ children }: ParentProps) {
  const navigate = useNavigate();
  const [activeNav, setActiveNav] = createSignal(1);
  const [sidebarOpened, setSidebarOpened] = createSignal(true);
  const [profilePopover, setProfilePopover] = createSignal(false);

  const auth = useContext(AuthContext)
  const sooner = useContext(SoonerContext)

  const [data, { refetch }] = createResource(GetStudentData);

  // const location = useLocation()

  createEffect(() => {
    if (data.error) {
      navigate("/");
    }
  }, [data.error]);

  createEffect(() => {
    if (data() == false) {
      navigate("/");
    }
  }, [data]);
  
  onMount(() => {
    refetch();
  });
  
  
  const handleLogout = async () => {
    sooner?.show("Information", "Logging out...")
    
    const status = await auth?.logout()
    
    
    if (status) {
      sooner?.show("Information", "Berhasil logout")
      navigate("/");
      return
    }
    
    sooner?.show("Information", "Gagal logout")

  }


  // TODO: Add listener to change the state of the side-nav depends on the current path
  // createEffect(() => {
  //   switch(location.pathname) {
  //     case "/student/dashboard":
  //       setActiveNav(1)
  //       return
  //   }
  // }, [location]);

  return (
    <div class="flex bg-white min-h-screen w-full">
      <Sidebar
        isOpened={() => sidebarOpened()}
        institutionLogoHref="https://i0.wp.com/www.uim.ac.id/uimv2/wp-content/uploads/2020/10/Ico.png?resize=200%2C200&ssl=1"
        institutionName="Universitas Islam Madura"
      >
        <SidebarNav
          id={1}
          title="Dashboard"
          isOpened={() => activeNav() == 1}
          isOpenSetter={setActiveNav}
        />
        <SidebarNav
          id={2}
          title="Biodata"
          isOpened={() => activeNav() == 2}
          isOpenSetter={setActiveNav}
        >
          <p>Data Diri</p>
          <p>Data Orang Tua</p>
          <p>Data Pendidikan</p>
          <p>Data Alamat</p>
          <p>Dokumen</p>
        </SidebarNav>
        <SidebarNav
          id={3}
          title="Akademik"
          isOpened={() => activeNav() == 3}
          isOpenSetter={setActiveNav}
        >
          <p>Jadwal Perkuliahan</p>
          <p>Kartu Rencana Studi (KRS)</p>
          <p>Kartu Hasil Studi (KHS)</p>
          <p>Kartu Tanda Mahasiswa (KTM)</p>
          <p>E-Kusioner</p>
        </SidebarNav>
      </Sidebar>
      <div class="w-screen">
        <header class="w-full flex items-center justify-between bg-background-700 px-4 py-6">
          <button
            on:click={() => setSidebarOpened((prev) => !prev)}
            class="w-max flex justify-center items-center hover:cursor-pointer"
          >
            <Icon icon="ci:hamburger-md" class="text-[32px]" />
          </button>
          {/* profile section */}
          <div class="flex items-center gap-x-4 relative">
            <div class="space-y-1 flex flex-col items-end">
              <Show
                when={data()}
                fallback={
                  <div class="h-[10px] w-[100px] bg-gray-600 animate-pulse "></div>
                }
              >
                <h4 class="font-medium">{data().name}</h4>
              </Show>
              <Show
                when={data()}
                fallback={
                  <div class="h-[10px] w-[100px] bg-gray-600 animate-pulse "></div>
                }
              >
                <span class="text-sm text-gray-400 capitalize">{data().user.role_name}</span>
              </Show>
            </div>

            {/* photo section */}
            <button
              on:click={() => setProfilePopover((prev) => !prev)}
              class="flex gap-x-2 items-center hover:cursor-pointer"
            >
              <div class="size-[50px] rounded-full overflow-hidden">
                <img
                  src="https://images.unsplash.com/photo-1500648767791-00dcc994a43e?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
                  alt="profile-photo"
                  class="w-full"
                />
              </div>
              <Icon
                icon="ion:chevron-down-sharp"
                class={`size-[18px] text-white`}
              />
            </button>
            {/* photo section */}
            {profilePopover() && (
              <button
                on:click={() => setProfilePopover(false)}
                class="fixed min-w-screen min-h-screen bg-black/0 z-20 top-0 left-0"
              ></button>
            )}

            {/* profile popover section */}
            <div
              class={`absolute z-30 right-0 -bottom-36 w-[250px] rounded-md bg-white p-4 text-center space-y-4 ${
                profilePopover()
                  ? "opacity-100 transition-opacity duration-300 ease-out"
                  : "opacity-0 transition-opacity duration-300 ease-in"
              }`}
            >
              <A href="/" class="text-black block hover:underline">
                Akun anda
              </A>
              <A href="/" class="text-black block hover:underline">
                Pengaturan
              </A>
              <button on:click={handleLogout} class="text-red-500 block hover:underline text-center w-full hover:cursor-pointer">
                Keluar
              </button>
            </div>
            {/* profile popover section */}
          </div>
          {/* profile section */}
        </header>
        {children}
      </div>
    </div>
  );
}
