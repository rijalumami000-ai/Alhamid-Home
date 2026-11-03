<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { EducationUnit } from '$lib/types';
  import { Plus, Edit2, Trash2, X, ExternalLink, GraduationCap } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let units = $state<EducationUnit[]>([]);
  let loading = $state(true);

  let isModalOpen = $state(false);
  let editingUnit = $state<EducationUnit | null>(null);

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let unitToDelete = $state<EducationUnit | null>(null);
  let isDeleting = $state(false);

  let form = $state({
    name: '',
    slug: '',
    type: 'formal' as 'formal' | 'non_formal',
    tagline: '',
    hero_image: '',
    badge: '',
    short_desc: '',
    full_desc: '',
    curriculum: '',
    requirements: '',
    facilities_list: '',
    headmaster_name: '',
    order: 0,
    is_active: true,
  });

  async function loadData() {
    loading = true;
    try {
      units = await api.admin.getEducationUnits();
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  function openCreate() {
    editingUnit = null;
    form = {
      name: '',
      slug: '',
      type: 'formal',
      tagline: '',
      hero_image: 'https://images.unsplash.com/photo-1523240795612-9a054b0db644?auto=format&fit=crop&w=1200&q=80',
      badge: 'Terakreditasi A',
      short_desc: '',
      full_desc: '',
      curriculum: '',
      requirements: '',
      facilities_list: '',
      headmaster_name: '',
      order: units.length + 1,
      is_active: true,
    };
    isModalOpen = true;
  }

  function openEdit(u: EducationUnit) {
    editingUnit = u;
    form = {
      name: u.name,
      slug: u.slug,
      type: u.type,
      tagline: u.tagline,
      hero_image: u.hero_image,
      badge: u.badge,
      short_desc: u.short_desc,
      full_desc: u.full_desc,
      curriculum: u.curriculum,
      requirements: u.requirements,
      facilities_list: u.facilities_list,
      headmaster_name: u.headmaster_name,
      order: u.order,
      is_active: u.is_active,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editingUnit) {
        await api.admin.updateEducationUnit(editingUnit.id, form);
      } else {
        await api.admin.createEducationUnit(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan unit pendidikan');
    }
  }

  function promptDelete(u: EducationUnit) {
    unitToDelete = u;
    isDeleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!unitToDelete) return;
    isDeleting = true;
    try {
      await api.admin.deleteEducationUnit(unitToDelete.id);
      isDeleteModalOpen = false;
      unitToDelete = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus unit pendidikan');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Unit Pendidikan | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Unit Pendidikan (Formal & Non-Formal)
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola Madrasah Diniyah, PMQ, MTs, MA, dan MAI Cintamulya.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
    >
      <Plus class="w-4 h-4" /> Tambah Unit Pendidikan
    </button>
  </div>

  <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs">
        <thead class="bg-slate-50 border-b border-slate-200 text-slate-500 font-bold uppercase tracking-wider text-[10px]">
          <tr>
            <th class="p-4">Nama Unit</th>
            <th class="p-4">Tipe</th>
            <th class="p-4">Badge / Akreditasi</th>
            <th class="p-4">Kepala Madrasah</th>
            <th class="p-4">Status</th>
            <th class="p-4 text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-100">
          {#each units as u}
            <tr class="hover:bg-slate-50/70 transition">
              <td class="p-4 flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-slate-100 overflow-hidden shrink-0 border border-slate-200">
                  {#if u.hero_image}
                    <img src={u.hero_image} alt={u.name} class="w-full h-full object-cover" />
                  {:else}
                    <div class="w-full h-full flex items-center justify-center text-teal-700 font-bold">
                      <GraduationCap class="w-5 h-5" />
                    </div>
                  {/if}
                </div>
                <div>
                  <div class="font-bold text-slate-900">{u.name}</div>
                  <div class="text-[11px] text-slate-400 font-mono">/{u.slug}</div>
                </div>
              </td>
              <td class="p-4">
                <span class="px-2 py-0.5 rounded font-semibold text-[10px] {u.type === 'formal' ? 'bg-blue-50 text-blue-700' : 'bg-amber-50 text-amber-700'}">
                  {u.type === 'formal' ? 'Formal (Kemenag)' : 'Non-Formal (Pesantren)'}
                </span>
              </td>
              <td class="p-4 font-medium text-slate-700">{u.badge || '-'}</td>
              <td class="p-4 text-slate-600">{u.headmaster_name || '-'}</td>
              <td class="p-4">
                <span class="px-2 py-0.5 rounded-full text-[10px] font-bold {u.is_active ? 'bg-emerald-100 text-emerald-800' : 'bg-slate-100 text-slate-600'}">
                  {u.is_active ? 'Aktif' : 'Nonaktif'}
                </span>
              </td>
              <td class="p-4 text-right">
                <div class="inline-flex items-center gap-1">
                  <a
                    href="/pendidikan/{u.slug}"
                    target="_blank"
                    class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
                    title="Pratinjau"
                  >
                    <ExternalLink class="w-4 h-4" />
                  </a>
                  <button
                    onclick={() => openEdit(u)}
                    class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
                    title="Edit Unit"
                  >
                    <Edit2 class="w-4 h-4" />
                  </button>
                  <button
                    onclick={() => promptDelete(u)}
                    class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
                    title="Hapus Unit"
                  >
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-3xl rounded-3xl shadow-2xl border border-slate-200 overflow-hidden flex flex-col max-h-[90vh]">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between bg-slate-50/50">
        <h2 class="font-display font-bold text-xl text-slate-900">
          {editingUnit ? 'Edit Unit Pendidikan' : 'Tambah Unit Pendidikan'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 overflow-y-auto space-y-4 flex-1">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1">
            <label for="u_name" class="text-xs font-bold text-slate-700 uppercase">Nama Unit *</label>
            <input
              id="u_name"
              type="text"
              required
              bind:value={form.name}
              placeholder="Contoh: MA Cintamulya"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="u_type" class="text-xs font-bold text-slate-700 uppercase">Tipe Pendidikan *</label>
            <select
              id="u_type"
              bind:value={form.type}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
            >
              <option value="formal">Pendidikan Formal (MTs, MA, MAI)</option>
              <option value="non_formal">Pendidikan Non-Formal (Diniyah, PMQ)</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1">
            <label for="u_badge" class="text-xs font-bold text-slate-700 uppercase">Badge / Akreditasi</label>
            <input
              id="u_badge"
              type="text"
              bind:value={form.badge}
              placeholder="Akreditasi A Unggul"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="u_head" class="text-xs font-bold text-slate-700 uppercase">Nama Kepala Madrasah / Pengasuh</label>
            <input
              id="u_head"
              type="text"
              bind:value={form.headmaster_name}
              placeholder="Nama dan gelar..."
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <div class="space-y-1">
          <label for="u_tagline" class="text-xs font-bold text-slate-700 uppercase">Tagline</label>
          <input
            id="u_tagline"
            type="text"
            bind:value={form.tagline}
            placeholder="Slogan atau tagline unit..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <DriveImageInput
          id="u_img"
          label="Hero Image / Foto Sampul Unit"
          bind:value={form.hero_image}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="u_short" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Singkat (Ringkasan)</label>
          <textarea
            id="u_short"
            rows="2"
            bind:value={form.short_desc}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="u_full" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Lengkap / Profil Unit</label>
          <textarea
            id="u_full"
            rows="4"
            bind:value={form.full_desc}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="u_curr" class="text-xs font-bold text-slate-700 uppercase">Kurikulum & Kitab yang Dikaji</label>
          <textarea
            id="u_curr"
            rows="4"
            bind:value={form.curriculum}
            placeholder="Daftar jenjang, kurikulum, dan kitab..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="u_req" class="text-xs font-bold text-slate-700 uppercase">Persyaratan Pendaftaran</label>
          <textarea
            id="u_req"
            rows="3"
            bind:value={form.requirements}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="pt-4 border-t border-slate-100 flex justify-end gap-3">
          <button
            type="button"
            onclick={() => (isModalOpen = false)}
            class="px-5 py-2.5 rounded-xl border border-slate-200 text-slate-600 text-xs font-bold"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-6 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Unit
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title="Hapus Unit Pendidikan"
  itemName={unitToDelete?.name}
  message="Unit pendidikan ini akan dihapus dari daftar akademik website."
  {isDeleting}
  onConfirm={confirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    unitToDelete = null;
  }}
/>
