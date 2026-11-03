<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Caregiver } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Users, BookOpen } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let list = $state<Caregiver[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Caregiver | null>(null);

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let caregiverToDelete = $state<Caregiver | null>(null);
  let isDeleting = $state(false);

  let form = $state({
    name: '',
    role_title: '',
    photo_url: '',
    bio: '',
    sanad_keilmuan: '',
    order: 0,
    is_active: true,
  });

  async function loadData() {
    try {
      list = await api.admin.getCaregivers();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      name: '',
      role_title: 'Dewan Masyayikh',
      photo_url: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=600&q=80',
      bio: '',
      sanad_keilmuan: '',
      order: list.length + 1,
      is_active: true,
    };
    isModalOpen = true;
  }

  function openEdit(item: Caregiver) {
    editing = item;
    form = {
      name: item.name,
      role_title: item.role_title,
      photo_url: item.photo_url,
      bio: item.bio,
      sanad_keilmuan: item.sanad_keilmuan,
      order: item.order,
      is_active: item.is_active,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateCaregiver(editing.id, form);
      } else {
        await api.admin.createCaregiver(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan data pengasuh');
    }
  }

  function promptDelete(item: Caregiver) {
    caregiverToDelete = item;
    isDeleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!caregiverToDelete) return;
    isDeleting = true;
    try {
      await api.admin.deleteCaregiver(caregiverToDelete.id);
      isDeleteModalOpen = false;
      caregiverToDelete = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus data');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Dewan Pengasuh | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Dewan Pengasuh & Masyayikh
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola profil Kyai, Masyayikh, dan pimpinan pesantren beserta silsilah sanad keilmuan.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
    >
      <Plus class="w-4 h-4" /> Tambah Dewan Pengasuh
    </button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between">
        <div class="p-6 space-y-4">
          <div class="flex items-start gap-4">
            <div class="w-20 h-20 rounded-2xl overflow-hidden bg-slate-900 shrink-0 border border-slate-200">
              <img src={item.photo_url} alt={item.name} class="w-full h-full object-cover" />
            </div>
            <div class="space-y-1 min-w-0">
              <span class="px-2 py-0.5 rounded text-[10px] font-bold uppercase {item.is_active ? 'bg-teal-100 text-teal-800' : 'bg-slate-200 text-slate-600'}">
                {item.is_active ? 'Aktif' : 'Nonaktif'}
              </span>
              <h2 class="font-display font-bold text-base text-slate-900 truncate">{item.name}</h2>
              <p class="text-xs text-teal-700 font-semibold">{item.role_title}</p>
            </div>
          </div>

          {#if item.bio}
            <p class="text-xs text-slate-500 line-clamp-3 leading-relaxed">{item.bio}</p>
          {/if}

          {#if item.sanad_keilmuan}
            <div class="p-2.5 rounded-xl bg-slate-50 border border-slate-100 text-[11px] text-slate-600 space-y-1">
              <div class="font-bold text-slate-800 flex items-center gap-1">
                <BookOpen class="w-3.5 h-3.5 text-teal-600" /> Sanad Keilmuan:
              </div>
              <p class="line-clamp-2">{item.sanad_keilmuan}</p>
            </div>
          {/if}
        </div>

        <div class="p-4 bg-slate-50/50 border-t border-slate-100 flex items-center justify-between">
          <span class="text-xs text-slate-400 font-mono">Urutan #{item.order}</span>
          <div class="flex items-center gap-1">
            <button
              onclick={() => openEdit(item)}
              class="p-2 rounded-xl text-slate-600 hover:text-teal-700 hover:bg-slate-100 transition"
              title="Edit Data"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => promptDelete(item)}
              class="p-2 rounded-xl text-slate-600 hover:text-rose-600 hover:bg-rose-50 transition"
              title="Hapus Data"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-lg rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editing ? 'Edit Dewan Pengasuh' : 'Tambah Dewan Pengasuh'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="c_name" class="text-xs font-bold text-slate-700 uppercase">Nama Lengkap & Gelar *</label>
          <input
            id="c_name"
            type="text"
            required
            bind:value={form.name}
            placeholder="K.H. ..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="c_role" class="text-xs font-bold text-slate-700 uppercase">Jabatan / Peran *</label>
          <input
            id="c_role"
            type="text"
            required
            bind:value={form.role_title}
            placeholder="Pengasuh Utama / Majelis Masyayikh"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <DriveImageInput
          id="c_photo"
          label="Foto Pengasuh / Masyayikh"
          bind:value={form.photo_url}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="c_bio" class="text-xs font-bold text-slate-700 uppercase">Biografi Singkat</label>
          <textarea
            id="c_bio"
            rows="3"
            bind:value={form.bio}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="c_sanad" class="text-xs font-bold text-slate-700 uppercase">Sanad Keilmuan</label>
          <textarea
            id="c_sanad"
            rows="2"
            bind:value={form.sanad_keilmuan}
            placeholder="Muttashil hingga..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="c_order" class="text-xs font-bold text-slate-700 uppercase">Urutan</label>
            <input
              id="c_order"
              type="number"
              bind:value={form.order}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="flex items-center gap-2 pt-6">
            <input id="c_active" type="checkbox" bind:checked={form.is_active} class="w-4 h-4 text-teal-600 rounded" />
            <label for="c_active" class="text-xs font-bold text-slate-700">Aktif Tampil</label>
          </div>
        </div>

        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button
            type="button"
            onclick={() => (isModalOpen = false)}
            class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Data
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title="Hapus Dewan Pengasuh"
  itemName={caregiverToDelete?.name}
  message="Profil dewan pengasuh ini akan dihapus dari sistem dan halaman profil pesantren."
  {isDeleting}
  onConfirm={confirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    caregiverToDelete = null;
  }}
/>
