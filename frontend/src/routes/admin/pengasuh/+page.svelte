<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Caregiver } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Users, BookOpen } from 'lucide-svelte';

  let list = $state<Caregiver[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Caregiver | null>(null);
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

  async function handleDelete(id: number) {
    if (confirm('Hapus data pengasuh ini?')) {
      try {
        await api.admin.deleteCaregiver(id);
        await loadData();
      } catch (err: any) {
        alert(err.message || 'Gagal menghapus');
      }
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
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Dewan Pengasuh & Masyayikh
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola profil Kyai, Masyayikh, dan pimpinan pesantren beserta silsilah sanad keilmuan.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
    >
      <Plus class="w-4 h-4" /> Tambah Pengasuh
    </button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
    {#each list as c}
      <div class="bg-white rounded-3xl border border-slate-200 p-6 flex flex-col justify-between shadow-sm space-y-4">
        <div class="flex items-start gap-4">
          <img
            src={c.photo_url}
            alt={c.name}
            class="w-16 h-16 rounded-2xl object-cover border border-slate-200 shrink-0"
          />
          <div class="space-y-1">
            <span class="text-[11px] font-bold text-emerald-700 uppercase">{c.role_title}</span>
            <h2 class="font-bold text-base text-slate-900">{c.name}</h2>
            <p class="text-xs text-slate-500 line-clamp-2">{c.bio}</p>
          </div>
        </div>

        {#if c.sanad_keilmuan}
          <div class="bg-emerald-50/70 p-3 rounded-xl border border-emerald-100 text-xs text-slate-600">
            <span class="font-bold text-emerald-800">Sanad:</span> {c.sanad_keilmuan}
          </div>
        {/if}

        <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
          <span class="text-xs text-slate-400 font-mono">Urutan: #{c.order}</span>
          <div class="space-x-1">
            <button
              onclick={() => openEdit(c)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-emerald-700 hover:bg-emerald-50"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => handleDelete(c.id)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-red-600 hover:bg-red-50"
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
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">
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

        <div class="space-y-1">
          <label for="c_photo" class="text-xs font-bold text-slate-700 uppercase">URL Foto</label>
          <input
            id="c_photo"
            type="url"
            bind:value={form.photo_url}
            placeholder="https://..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

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
            class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md"
          >
            Simpan Data
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

