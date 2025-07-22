import { MetaProvider, Title } from "@solidjs/meta";
import { createSignal } from "solid-js";
import { $ZodIssue } from "zod/v4/core";
import Breadcrumps from "~/components/atoms/breadcrumps";
import FormGroup from "~/components/atoms/forms/form-group";
import InputForm from "~/components/atoms/forms/input-form";
import DashboardLayout from "~/components/layouts/DashboardLayout";

export default function PersonalDataPage() {

  const [errors, setErrors] = createSignal<$ZodIssue[]>([])

  const isInvalid = (fieldName: string, grabMessegeOnly?: boolean) => {
    const expectedField = errors().filter(fieldIssue => fieldIssue.path[0] == fieldName)


    if (grabMessegeOnly && expectedField.length >= 1) {
        return expectedField[0].message
    }

    if (expectedField.length >= 1) return true

    return false
  }

  return (
    <DashboardLayout>
      <MetaProvider>
        <Title>Zenith - Data Diri Mahasiswa</Title>
      </MetaProvider>

      <main class="text-left mx-auto text-gray-700 p-8 bg-background-800 min-h-screen">
        <Breadcrumps paths={[{label: "Biodata", href: "/student/biodata/data-diri"}]} />
        <div class="mt-4">
          <FormGroup title="Data diri mahasiswa">
             <InputForm invalidMessege={() => isInvalid("student_name", true)} isInvalid={() => isInvalid("student_name")} title="Nama mahasiswa" name="student_name" type="email" />
          </FormGroup>
        </div>
      </main>
    </DashboardLayout>
  );
}
